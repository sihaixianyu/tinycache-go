package lru

import "tinycache-go/list"

type Cache[T list.Sized] struct {
	// Memory control
	maxBytes  uint64
	usedBytes uint64
	// Core data stucture
	nodeList *list.List[T]
	nodeMap  map[string]*list.Node[T]
	// Action to be performed when an entry is cleared
	OnCleared func(key string, Val T)
}

func NewCache[T list.Sized](maxBytes uint64, onCleared func(key string, Val T)) *Cache[T] {
	return &Cache[T]{
		maxBytes:  maxBytes,
		usedBytes: 0,
		nodeList:  list.New[T](),
		nodeMap:   make(map[string]*list.Node[T]),
		OnCleared: onCleared,
	}
}

func (c *Cache[T]) Get(key string) (val T, ok bool) {
	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		return node.Val, true
	}

	return
}

func (c *Cache[T]) Add(key string, val T) {
	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		c.usedBytes += val.Size() - node.Val.Size()
		node.Val = val
	} else {
		node := &list.Node[T]{Key: key, Val: val}
		c.nodeList.PushFront(node)
		c.nodeMap[key] = node
	}

	for c.maxBytes != 0 && c.maxBytes < c.usedBytes {
		c.nodeList.PopBack()
	}
}

func (c *Cache[T]) RemoveOldest() {
	node := c.nodeList.PopBack()
	if node != nil {
		key, val := node.Key, node.Val
		delete(c.nodeMap, node.Key)
		c.usedBytes -= uint64(len(key)) + val.Size()

		if c.OnCleared != nil {
			c.OnCleared(key, val)
		}
	}
}

// Return length of nodeList, which also means the number of elements in cache
func (c *Cache[T]) Len() uint {
	return c.nodeList.Len()
}

// Return the used bytes of cache
func (c *Cache[T]) Size() uint64 {
	return c.usedBytes
}

// Return the maximum bytes of cache
func (c *Cache[T]) Cap() uint64 {
	return c.maxBytes
}
