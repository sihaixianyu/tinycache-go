package list

import (
	"fmt"
)

type Sized interface {
	Size() uint64
}

type Node[T Sized] struct {
	prev *Node[T]
	next *Node[T]
	list *List[T]

	Key string
	Val  T
}

type List[T Sized] struct {
	dummyHead *Node[T]
	dummyTail *Node[T]
	len       uint
}

func New[T Sized]() *List[T] {
	list := &List[T]{
		dummyHead: &Node[T]{},
		dummyTail: &Node[T]{},
		len:       0,
	}

	list.dummyHead.next = list.dummyTail
	list.dummyTail.prev = list.dummyHead

	list.dummyHead.list = list
	list.dummyTail.list = list

	return list
}

func (l *List[T]) PushFront(node *Node[T]) *Node[T] {
	head := l.dummyHead.next

	l.dummyHead.next = node
	node.prev = l.dummyHead
	node.next = head
	head.prev = node

	node.list = l
	l.len += 1

	return node
}

func (l *List[T]) PopBack() *Node[T] {
	if l.dummyTail.prev == l.dummyHead {
		return nil
	}

	tar := l.dummyTail.prev
	tar.prev.next = l.dummyTail
	l.dummyTail = tar.prev

	tar.list = nil
	l.len -= 1

	return tar
}

func (l *List[T]) MoveToFront(node *Node[T]) {
	if node.list != l {
		return
	}

	l.remove(node)
	l.PushFront(node)
}

func (l *List[T]) Debug() {
	for curr := l.dummyHead.next; curr != nil; curr = curr.next {
		fmt.Printf("%v ", curr)
	}
	fmt.Println()
}

func (l *List[T]) remove(node *Node[T]) {
	if node.list != l {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil
	l.len -= 1
}

func (l *List[T]) Len() uint {
	return l.len
}