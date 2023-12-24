package linkedlist

import "errors"

// Define the List and Element types here.
type Element struct {
	value int
	next  *Element
}

type List struct {
	head *Element
	len  int
}

func New(elements []int) *List {
	result := &List{}
	for _, v := range elements {
		result.Push(v)
	}
	return result
}

func (l *List) Size() int {
	if l == nil {
		return 0
	}
	return l.len
}

func (l *List) Push(element int) {
	l.head = &Element{value: element, next: l.head}
	l.len++
}

func (l *List) Pop() (int, error) {
	if l.len == 0 {
		return -1, errors.New("Linked list is empty")
	} else {
		result := l.head.value
		l.head = l.head.next
		l.len--
		return result, nil
	}
}

func (l *List) Array() []int {
	var result []int
	for l.head != nil {
		e, _ := l.Pop()
		result = append([]int{e}, result...)
	}
	return result
}

func (l *List) Reverse() *List {
	revList := New([]int{})
	for l.head != nil {
		e, _ := l.Pop()
		revList.Push(e)
	}
	l = revList
	return l
}
