package lru

import (
	"fmt"
	"testing"
	"tinycache-go/util"

	"github.com/stretchr/testify/assert"
)

type Space struct {
	size int
}

func (s Space) Size() int {
	return s.size
}

func TestLRUCache(t *testing.T) {
	c := New(3, func(key string, val Space) {
		fmt.Printf("clear entry: key=\"%s\" size=%dB\n", key, val.Size())
	})

	space1 := Space{size: 1}
	space2 := Space{size: 2}
	space3 := Space{size: 3}

	c.Put("space1", space1)
	c.Put("space2", space2)
	c.Put("space3", space3)
	assert.Equal(t, 3, c.Cap())
	assert.Equal(t, 3, c.Len())
	util.Debug(c)

	v, ok := c.Get("space1")
	assert.Equal(t, true, ok)
	assert.Equal(t, space1, v)
	util.Debug(c)

	space4 := Space{size: 4}
	c.Put("space4", space4)
	assert.Equal(t, 3, c.Cap())
	assert.Equal(t, 3, c.Len())
	util.Debug(c)

	assert.Equal(t, 3, c.Cap())
	assert.Equal(t, 3, c.Len())

	c.RemoveLast()
	v, ok = c.Get("space2")
	assert.Equal(t, false, ok)
	assert.Equal(t, v, Space{})
	util.Debug(c)
}
