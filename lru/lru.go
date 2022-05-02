package lru

import (
	"fmt"
	"strings"
	"tinycache-go/list"
	"tinycache-go/util"
)

type Sized interface {
	Size() int
}

type entry[K comparable, V Sized] struct {
	Key K
	Val V
}

type LRU[K comparable, V Sized] struct {
	cap int
	len int

	nodeMap  map[K]*list.Node[*entry[K, V]]
	nodeList *list.List[*entry[K, V]]

	// Action to be performed when an entry is cleared
	OnCleared func(key K, val V)
}

// * If invoker set maxBytes to 0, the available cache is math.MaxInt
func New[K comparable, V Sized](cap int, onCleared func(key K, val V)) *LRU[K, V] {
	return &LRU[K, V]{
		cap:       cap,
		nodeList:  list.New[*entry[K, V]](),
		nodeMap:   make(map[K]*list.Node[*entry[K, V]]),
		OnCleared: onCleared,
	}
}

func (c *LRU[K, V]) Cap() int {
	return c.cap
}

func (c *LRU[K, V]) Len() int {
	return c.len
}

func (c *LRU[K, V]) Put(key K, val V) {
	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		node.Elem.Val = val
	} else {
		entry := &entry[K, V]{Key: key, Val: val}
		newNode := c.nodeList.PushFront(entry)
		c.nodeMap[key] = newNode
		c.len += 1
	}

	for c.cap != 0 && c.cap < c.len {
		c.RemoveLast()
	}
}

func (c *LRU[K, V]) Get(key K) (val V, ok bool) {
	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		return node.Elem.Val, true
	}

	return
}

func (c *LRU[K, V]) RemoveLast() {
	node := c.nodeList.PopBack()
	if node != nil {
		key, val := node.Elem.Key, node.Elem.Val
		delete(c.nodeMap, node.Elem.Key)
		c.len -= 1

		if c.OnCleared != nil {
			c.OnCleared(key, val)
		}
	}
}

func (c *LRU[K, V]) Format(level int) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Cache: %p {\n", c))

	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("Capacity: %d\n", c.cap))
	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("Length: %d\n", c.len))
	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("nodeMap: %v\n", c.nodeMap))
	util.InsertTab(&builder, level+1)
	builder.WriteString(fmt.Sprintf("nodeList: %s\n", c.nodeList.Format(1)))

	util.InsertTab(&builder, level)
	builder.WriteString("}")

	return builder.String()
}
