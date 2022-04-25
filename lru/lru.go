package lru

import (
	"fmt"
	"strings"
	"tinycache-go/list"
	"tinycache-go/util"
)

type Sized interface {
	Size() uint64
}

type Entry struct {
	Key string
	Val Sized
}

type Cache struct {
	// Memory control
	maxBytes  uint64
	usedBytes uint64
	// Core data stucture
	nodeList *list.List[*Entry]
	nodeMap  map[string]*list.Node[*Entry]
	// Action to be performed when an entry is cleared
	OnCleared func(key string, val Sized)
}

// * If invoker set maxBytes to 0, the available cache is infinite
func New(maxBytes uint64, onCleared func(key string, val Sized)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		usedBytes: 0,
		nodeList:  list.New[*Entry](),
		nodeMap:   make(map[string]*list.Node[*Entry]),
		OnCleared: onCleared,
	}
}

// Return the number of bytes used in current cache
func (c *Cache) UsedBytes() uint64 {
	return c.usedBytes
}

// Return the number of maximum bytes allowed in current cache
func (c *Cache) MaxBytes() uint64 {
	return c.maxBytes
}

func (c *Cache) Get(key string) (val Sized, ok bool) {
	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		return node.Elem.Val, true
	}

	return
}

func (c *Cache) Add(key string, val Sized) {
	if node, ok := c.nodeMap[key]; ok {
		c.nodeList.MoveToFront(node)
		node.Elem.Val = val
		c.usedBytes += val.Size() - node.Elem.Val.Size()
	} else {
		entry := &Entry{Key: key, Val: val}
		newNode := c.nodeList.PushFront(entry)
		c.nodeMap[key] = newNode
		c.usedBytes += uint64(len(key)) + val.Size()
	}

	for c.maxBytes != 0 && c.maxBytes < c.usedBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) RemoveOldest() {
	node := c.nodeList.PopBack()
	if node != nil {
		key, val := node.Elem.Key, node.Elem.Val
		delete(c.nodeMap, node.Elem.Key)
		c.usedBytes -= uint64(len(key)) + val.Size()

		if c.OnCleared != nil {
			c.OnCleared(key, val)
		}
	}
}

func (c *Cache) Format(level int) string {
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
