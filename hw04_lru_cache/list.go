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
	lenlist   int
	frontlist *ListItem
	backlist  *ListItem
}

func (list *list) Len() int {
	return list.lenlist
}

func (list *list) Front() *ListItem {
	return list.frontlist
}

func (list *list) Back() *ListItem {
	return list.backlist
}

func (list *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{
		Value: v,
	}

	if list.frontlist != nil {
		list.insertBefore(list.frontlist, &newItem)
	} else {
		list.frontlist = &newItem
		list.backlist = &newItem
		newItem.Next = nil
		newItem.Prev = nil
	}

	list.lenlist++

	return &newItem
}

func (list *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{
		Value: v,
	}
	if list.backlist != nil {
		list.insertAfter(list.backlist, &newItem)
		list.lenlist++
		return &newItem
	} else {
		return list.PushFront(v)
	}
}

func (list *list) Remove(i *ListItem) {

}
func NewList() List {
	return new(list)
}
