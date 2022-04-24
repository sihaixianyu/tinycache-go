package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Space struct {
	size uint64
}

func (s Space) Size() uint64 {
	return s.size
}

func TestList(t *testing.T) {
	l := New[Space]()
	assert.Equal(t, 0, l.len)

	node1 := &Node[Space]{
		Val: Space{size: 1},
	}
	l.PushFront(node1)

	node2 := &Node[Space]{
		Val: Space{size: 2},
	}
	l.PushFront(node2)

	node3 := &Node[Space]{
		Val: Space{size: 4},
	}
	l.PushFront(node3)

	l.Debug()

	l.MoveToFront(node2)
	l.Debug()

	l.MoveToFront(node3)
	l.Debug()

	l.remove(node3)
	l.Debug()

	l.remove(node1)
	l.Debug()

	l.remove(node2)
	l.Debug()
}
