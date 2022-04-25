package lru

import (
	"fmt"
	"testing"
	"tinycache-go/util"

	"github.com/stretchr/testify/assert"
)

type Space struct {
	size uint64
}

func (s Space) Size() uint64 {
	return s.size
}

func TestCache(t *testing.T) {
	c := New(36, func(key string, val Sized) {
		fmt.Printf("clear entry: key=%s size=%d", key, val.Size())
	})

	space1 := Space{size: 6}
	space2 := Space{size: 6}
	space3 := Space{size: 6}

	c.Add("space1", space1)
	c.Add("space2", space2)
	c.Add("space3", space3)
	assert.Equal(t, uint64(36), c.MaxBytes())
	assert.Equal(t, uint64(36), c.UsedBytes())
	util.Debug(c)

	v, ok := c.Get("space1")
	assert.Equal(t, true, ok)
	assert.Equal(t, space1, v)
	util.Debug(c)

	space4 := Space{size: 6}
	c.Add("space4", space4)
	assert.Equal(t, uint64(36), c.MaxBytes())
	assert.Equal(t, uint64(36), c.UsedBytes())
	util.Debug(c)

	assert.Equal(t, uint64(36), c.MaxBytes())
	assert.Equal(t, uint64(36), c.UsedBytes())

	c.RemoveOldest()
	v, ok = c.Get("space2")
	assert.Equal(t, false, ok)
	assert.Equal(t, v, nil)
	util.Debug(c)
}
