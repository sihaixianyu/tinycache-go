package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := &Cache{}

	space1 := ByteSpace{'1'}
	cache.Put("k1", space1)

	space2 := ByteSpace{'2'}
	cache.Put("k2", space2)

	v1, _ := cache.Get("k1")
	assert.Equal(t, space1, v1)

	v2, _ := cache.Get("k2")
	assert.Equal(t, space2, v2)

	_, ok := cache.Get("k3")
	assert.False(t, ok)
}
