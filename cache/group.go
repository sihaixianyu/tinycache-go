package cache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	mainCache Cache
	getter    Getter
	picker    PeerPicker
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, maxBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}

	mu.Lock()
	defer mu.Unlock()

	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: NewCache(maxBytes),
	}

	groups[name] = g

	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()

	return g
}

func (g *Group) RegisterPicker(picker PeerPicker) {
	if g.picker != nil {
		panic("register peer picker more than once")
	}

	g.picker = picker
}

func (g *Group) Get(key string) (ByteSpace, error) {
	if key == "" {
		return ByteSpace{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.Get(key); ok {
		log.Println("[TinyCache] hit")
		return v, nil
	}

	return g.load(key)
}

func (g *Group) load(key string) (ByteSpace, error) {
	return g.getLocal(key)
}

func (g *Group) getLocal(key string) (ByteSpace, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteSpace{}, err
	}

	val := ByteSpace(bytes)
	g.populateCache(key, val)

	return val, nil
}

func (g *Group) populateCache(key string, value ByteSpace) {
	g.mainCache.Put(key, value)
}
