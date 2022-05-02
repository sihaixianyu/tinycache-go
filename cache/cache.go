package cache

import (
	"sync"
	"tinycache-go/lru"
)

type ByteSpace []byte

// Implementation of cache.Sized interface
func (c ByteSpace) Size() int {
	return len(c)
}

func (c ByteSpace) String() string {
	return string(c)
}

type Cache struct {
	mu                  sync.RWMutex
	lru                 *lru.LRU[string, ByteSpace]
	maxBytes, usedBytes int64
	nHit, nGet          int64
	nCleared            int64
}

func NewCache(maxBytes int64) Cache {
	return Cache{
		maxBytes: maxBytes,
	}
}

func (c *Cache) Cap() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.maxBytes
}

func (c *Cache) Size(key string) int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.usedBytes
}

func (c *Cache) Put(key string, val ByteSpace) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		c.lru = lru.New(0, func(key string, val ByteSpace) {
			c.usedBytes -= int64(len(key)) + int64(val.Size())
			c.nCleared += 1
		})
	}

	nBytes := int64(len(key)) + int64(val.Size())
	for c.usedBytes+nBytes > c.maxBytes {
		c.lru.RemoveLast()
	}

	c.lru.Put(key, val)
	c.usedBytes += nBytes
}

func (c *Cache) Get(key string) (value ByteSpace, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.nGet += 1
	if c.lru == nil {
		return
	}

	v, ok := c.lru.Get(key)
	if !ok {
		return
	}
	c.nHit += 1

	return v, true
}
