package lru

import (
	"fmt"
	"testing"
	"tinycache-go/util"

	"github.com/stretchr/testify/assert"
)

type Value struct {
	size int
}

func TestLRU(t *testing.T) {
	c := New(3, func(key string, val Value) {
		fmt.Printf("clear entry: key=\"%s\" val=\"%v\"\n", key, val)
	})

	space1 := Value{size: 1}
	space2 := Value{size: 2}
	space3 := Value{size: 3}

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

	space4 := Value{size: 4}
	c.Put("space4", space4)
	assert.Equal(t, 3, c.Cap())
	assert.Equal(t, 3, c.Len())
	util.Debug(c)

	assert.Equal(t, 3, c.Cap())
	assert.Equal(t, 3, c.Len())

	c.RemoveLast()
	v, ok = c.Get("space2")
	assert.Equal(t, false, ok)
	assert.Equal(t, v, Value{})
	util.Debug(c)
}
