package cache

import (
	"sync"
	"tinycache-go/cache"
	"tinycache-go/cache/lru"
)

type Getter interface {
	Get(key string) (cache.Sized, error)
}

type GetterFunc func(key string) (cache.Sized, error)

func (f GetterFunc) Get(key string) (cache.Sized, error) {
	return f(key)
}

type Group struct {
	name   string
	cache  cache.Cacher
	getter Getter
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, maxBytes int, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}

	mu.Lock()
	defer mu.Unlock()

	g := &Group{
		name:   name,
		getter: getter,
		cache:  lru.NewLRUCache(maxBytes, nil),
	}

	groups["name"] = g

	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()

	return g
}
