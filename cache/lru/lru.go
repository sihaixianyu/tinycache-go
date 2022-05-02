package lru

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"tinycache-go/cache"
	"tinycache-go/list"
	"tinycache-go/util"
)

type LRUCache struct {
	mu sync.Mutex

	maxBytes  int
	usedBytes int

	nodeList *list.List[*cache.Entry]
	nodeMap  map[string]*list.Node[*cache.Entry]

	// Action to be performed when an entry is cleared
	OnCleared func(key string, val cache.Sized)
}

// * If invoker set maxBytes to 0, the available cache is math.MaxUint64
func NewLRUCache(maxBytes int, onCleared func(key string, val cache.Sized)) *LRUCache {
	if maxBytes == 0 {
		maxBytes = math.MaxInt64
	}

	return &LRUCache{
		maxBytes:  maxBytes,
		usedBytes: 0,
		nodeList:  list.New[*cache.Entry](),
		nodeMap:   make(map[string]*list.Node[*cache.Entry]),
		OnCleared: onCleared,
	}
}

// Return the number of maximum bytes allowed in current cache
func (c *LRUCache) MaxBytes() int {
	return c.maxBytes
}

// Return the number of bytes used in current cache
func (c *LRUCache) UsedBytes() int {
	return c.usedBytes
}

func (c *LRUCache) Set(key string, val cache.Sized) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		node.Elem.Val = val
		c.usedBytes += val.Size() - node.Elem.Val.Size()
	} else {
		entry := &cache.Entry{Key: key, Val: val}
		newNode := c.nodeList.PushFront(entry)
		c.nodeMap[key] = newNode
		c.usedBytes += len(key) + val.Size()
	}

	for c.maxBytes != 0 && c.maxBytes < c.usedBytes {
		c.removeLast()
	}
}

func (c *LRUCache) Get(key string) (val cache.Sized, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		return node.Elem.Val, true
	}

	return
}

func (c *LRUCache) Format(level int) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Cache: %p {\n", c))

	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("UsedBytes: %d\n", c.usedBytes))
	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("MaxBytes: %d\n", c.maxBytes))
	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("nodeMap: %v\n", c.nodeMap))
	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("nodeList: %s\n", c.nodeList.Format(1)))

	util.InsertTab(&builder, level)
	builder.WriteString("}")

	return builder.String()
}

func (c *LRUCache) removeLast() {
	node := c.nodeList.PopBack()
	if node != nil {
		key, val := node.Elem.Key, node.Elem.Val
		delete(c.nodeMap, node.Elem.Key)
		c.usedBytes -= len(key) + val.Size()

		if c.OnCleared != nil {
			c.OnCleared(key, val)
		}
	}
}
