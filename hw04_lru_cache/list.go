package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.front = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	if l.len == 0 {
		return
	}
	l.len--
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{
		Value: v,
		Next:  l.front,
	}

	if item.Next != nil {
		l.front.Prev = &item
	}

	if l.back == nil {
		l.back = &item
	}

	l.len++
	l.front = &item
	return &item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{
		Value: v,
		Prev:  l.back,
	}

	if item.Prev == nil {
		l.front = &item
	} else {
		l.back.Next = &item
	}

	l.len++
	l.back = &item
	return &item
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i {
		return
	}

	if i.Next == nil {
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
		i.Next = l.front
		l.front.Prev = i
	}

	i.Prev = nil
	l.front = i
}
