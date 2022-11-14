package linkedlist

import "errors"

// ErrEmptyList represents error caused by empty linked list.
var ErrEmptyList = errors.New("empty list")

// Node is element of a linked list, which holds a value and pointers to the next and previous nodes.
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// List represents double linked list, which holds references to the first and last node.
type List struct {
	first *Node
	last  *Node
}

// NewList creates new linked list from variable number of interface inputs.
func NewList(args ...interface{}) *List {
	var result List
	for _, v := range args {
		result.PushBack(v)
	}
	return &result
}

// Next returns node preceding node in input.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns node following node in input.
func (n *Node) Prev() *Node {
	return n.prev
}

// PushFront inserts node at the beginning of the list.
func (l *List) PushFront(v interface{}) {
	node := &Node{
		Val:  v,
		next: l.first,
	}
	if l.first != nil {
		l.first.prev = node
	} else {
		l.last = node
	}
	l.first = node
}

// PushBack inserts node at the end of the list.
func (l *List) PushBack(v interface{}) {
	node := &Node{
		Val:  v,
		prev: l.last,
	}
	if l.last != nil {
		l.last.next = node
	} else {
		l.first = node
	}
	l.last = node
}

// PopFront removes node from the beginning of the list and returns it.
func (l *List) PopFront() (interface{}, error) {
	if l.first == nil {
		return nil, ErrEmptyList
	}
	result := l.first.Val
	if l.first == l.last {
		l.last = nil
	}
	l.first = l.first.next
	if l.first != nil {
		l.first.prev = nil
	}
	return result, nil
}

// PopBack removes node from the end of the list and returns it.
func (l *List) PopBack() (interface{}, error) {
	if l.last == nil {
		return nil, ErrEmptyList
	}
	result := l.last.Val
	if l.last == l.first {
		l.first = nil
	}
	l.last = l.last.prev
	if l.last != nil {
		l.last.next = nil
	}
	return result, nil
}

// Reverse reverses entered list ;-)
func (l *List) Reverse() {
	if l.first == nil || l.first.next == nil {
		return
	}
	node := l.first
	for node != nil {
		node.prev, node.next = node.next, node.prev
		node = node.prev
	}
	l.last, l.first = l.first, l.last
}

// First returns node from the beginning of the list.
func (l *List) First() *Node {
	return l.first
}

// Last returns node from the end of the list.
func (l *List) Last() *Node {
	return l.last
}
