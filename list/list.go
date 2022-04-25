package list

import (
	"fmt"
	"strings"
	"tinycache-go/util"
)

type Node[T any] struct {
	prev *Node[T]
	next *Node[T]
	list *List[T]

	Elem T
}

type List[T any] struct {
	dummyHead *Node[T]
	dummyTail *Node[T]
	len       uint
}

func New[T any]() *List[T] {
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

func (l List[T]) Len() uint {
	return l.len
}

func (l List[T]) Front() *Node[T] {
	if l.dummyHead.next == l.dummyTail {
		return nil
	}

	return l.dummyHead.next
}

func (l List[T]) Back() *Node[T] {
	if l.dummyHead.next == l.dummyTail {
		return nil
	}

	return l.dummyTail.prev
}

func (l *List[T]) PushFront(val T) *Node[T] {
	currHead := l.dummyHead.next
	newNode := &Node[T]{Elem: val}

	l.dummyHead.next = newNode
	currHead.prev = newNode

	newNode.prev = l.dummyHead
	newNode.next = currHead

	newNode.list = l
	l.len += 1

	return newNode
}

func (l *List[T]) PushBack(val T) *Node[T] {
	currTail := l.dummyTail.prev
	newNode := &Node[T]{Elem: val}

	l.dummyTail.prev = newNode
	currTail.next = newNode

	newNode.next = l.dummyTail
	newNode.prev = currTail

	newNode.list = l
	l.len += 1

	return newNode
}

func (l *List[T]) PopFront() *Node[T] {
	if l.dummyTail.prev == l.dummyHead {
		return nil
	}

	tar := l.dummyHead.next
	l.remove(tar)

	return tar
}

func (l *List[T]) PopBack() *Node[T] {
	if l.dummyTail.prev == l.dummyHead {
		return nil
	}

	tar := l.dummyTail.prev
	l.remove(tar)

	return tar
}

func (l *List[T]) MoveToFront(node *Node[T]) {
	if node.list != l {
		return
	}

	l.remove(node)
	l.PushFront(node.Elem)
}

func (l List[T]) Format(level int) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("list: %p {\n", &l))

	for curr := l.dummyHead.next; curr != l.dummyTail; curr = curr.next {
		util.InsertTab(&builder, level+1)
		builder.WriteString(fmt.Sprintf("node: %p {prev: %p, next: %p, Elem: %v}\n", curr, curr.prev, curr.next, curr.Elem))
	}

	util.InsertTab(&builder, level)
	builder.WriteString("}")

	return builder.String()
}

func (l *List[T]) remove(node *Node[T]) {
	if node.list != l {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	// Avoid memory leak
	node.prev = nil
	node.next = nil

	// No longer belong to this list
	node.list = nil
	l.len -= 1
}
