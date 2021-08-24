package hw04_lru_cache //nolint:golint,stylecheck

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

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.FirstItem
}

func (l *list) Back() *ListItem {
	return l.LastItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.Length == 0 {
		l.LastItem = item
	} else {
		item.Next = l.FirstItem
		l.FirstItem.Prev = item
	}
	l.FirstItem = item
	l.Length++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.Length == 0 {
		l.FirstItem = item
	} else {
		item.Prev = l.LastItem
		l.LastItem.Next = item
	}
	l.LastItem = item
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
		l.LastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.Length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
