package list

import (
	"testing"
	"tinycache-go/util"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	l := New[int]()
	assert.Equal(t, uint(0), l.len)

	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	util.Debug(l)

	assert.Equal(t, 3, l.Front().Elem)
	assert.Equal(t, 1, l.Back().Elem)

	l.MoveToFront(l.Back())
	assert.Equal(t, 1, l.Front().Elem)
	util.Debug(l)

	assert.Equal(t, 1, l.PopFront().Elem)
	util.Debug(l)

	assert.Equal(t, 2, l.PopBack().Elem)
	util.Debug(l)

	assert.Equal(t, 3, l.PopBack().Elem)
	util.Debug(l)
}
