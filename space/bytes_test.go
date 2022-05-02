package space

import (
	"testing"
	"tinycache-go/cache/lru"

	"github.com/stretchr/testify/assert"
)

func TestBytesSpace(t *testing.T) {
	cache := lru.NewLRUCache(0, nil)

	space1 := BytesSpace{'1'}
	cache.Set("k1", space1)

	space2 := BytesSpace{'2'}
	cache.Set("k2", space2)

	v1, _ := cache.Get("k1")
	assert.Equal(t, space1, v1)

	v2, _ := cache.Get("k2")
	assert.Equal(t, space2, v2)

	_, ok := cache.Get("k3")
	assert.False(t, ok)
}
