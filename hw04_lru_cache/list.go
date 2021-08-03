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
	Length    int
	FirstItem *ListItem
	LastItem  *ListItem
}

func (l list) Len() int {
	return l.Length
}

func (l list) Front() *ListItem {
	return l.FirstItem
}

func (l list) Back() *ListItem {
	return l.LastItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
	}
	if l.Length == 0 {
		l.FirstItem = item
		l.LastItem = item
	} else {
		nextItem := l.FirstItem
		l.FirstItem = item
		if nextItem != nil {
			nextItem.Prev = item
		}
		item.Next = nextItem
	}
	l.Length++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
	}
	if l.Length == 0 {
		l.FirstItem = item
		l.LastItem = item
	} else {
		prevItem := l.LastItem
		l.LastItem = item
		if prevItem != nil {
			prevItem.Next = item
		}
		item.Prev = prevItem
	}
	l.Length++
	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.FirstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.FirstItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.Length--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Length == 0 || i.Prev == nil {
		return
	}
	i.Prev.Next = i.Next
	if i.Next == nil {
		l.LastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.FirstItem.Prev = i
	i.Next = l.FirstItem
	i.Prev = nil
	l.FirstItem = i
}

func NewList() List {
	return new(list)
}
