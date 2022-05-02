package lru

import (
	"fmt"
	"testing"
	"tinycache-go/cache"
	"tinycache-go/util"

	"github.com/stretchr/testify/assert"
)

type Space struct {
	size int
}

func (s Space) Size() int {
	return s.size
}

func TestCache(t *testing.T) {
	c := NewLRUCache(36, func(key string, val cache.Sized) {
		fmt.Printf("clear entry: key=\"%s\" size=%dB\n", key, val.Size())
	})

	space1 := Space{size: 6}
	space2 := Space{size: 6}
	space3 := Space{size: 6}

	c.Set("space1", space1)
	c.Set("space2", space2)
	c.Set("space3", space3)
	assert.Equal(t, 36, c.MaxBytes())
	assert.Equal(t, 36, c.UsedBytes())
	util.Debug(c)

	v, ok := c.Get("space1")
	assert.Equal(t, true, ok)
	assert.Equal(t, space1, v)
	util.Debug(c)

	space4 := Space{size: 6}
	c.Set("space4", space4)
	assert.Equal(t, 36, c.MaxBytes())
	assert.Equal(t, 36, c.UsedBytes())
	util.Debug(c)

	assert.Equal(t, 36, c.MaxBytes())
	assert.Equal(t, 36, c.UsedBytes())

	c.removeLast()
	v, ok = c.Get("space2")
	assert.Equal(t, false, ok)
	assert.Equal(t, v, nil)
	util.Debug(c)
}
