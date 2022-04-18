package lru

import "container/list"

type Cache[T comparable] struct {
	// Memory control
	maxBytes  uint
	usedBytes uint
	// Core data stucture
	nodeList *list.List
	nodeMap  map[T]*list.Element
	// Action to be performed when an entry is cleared
	OnCleared func(key T, Val Sized)
}

type Sized interface {
	Size() uint
}

type entry[T comparable] struct {
	key T
	val Sized
}

func NewCache[T comparable](maxBytes uint, onCleared func(key T, Val Sized)) *Cache[T] {
	return &Cache[T]{
		maxBytes:  maxBytes,
		usedBytes: 0,
		nodeList:  list.New(),
		nodeMap:   make(map[T]*list.Element),
		OnCleared: onCleared,
	}
}

// TODO:
func Get[T comparable](key T) (val Sized, ok bool) {
	return nil, false
}
