package deque

import "github.com/ndrewnee/go-algods/ds/list"

var _ Deque = &ListDeque{}

type ListDeque struct {
	list *list.List
}

func NewListDeque() *ListDeque {
	return &ListDeque{list: list.New()}
}

func (d *ListDeque) PushBack(value interface{}) {
	_ = d.list.PushBack(value)
}

func (d *ListDeque) PushFront(value interface{}) {
	_ = d.list.PushFront(value)
}

func (d *ListDeque) PopBack() interface{} {
	tail := d.list.Tail()
	if tail == nil {
		return nil
	}
	d.list.Remove(tail)
	return tail.Value()
}

func (d *ListDeque) PopFront() interface{} {
	head := d.list.Head()
	if head == nil {
		return nil
	}
	d.list.Remove(head)
	return head.Value()
}

func (d *ListDeque) Front() interface{} {
	head := d.list.Head()
	if head == nil {
		return nil
	}
	return head.Value()
}

func (d *ListDeque) Back() interface{} {
	tail := d.list.Tail()
	if tail == nil {
		return nil
	}
	return tail.Value()
}

func (d *ListDeque) Size() int {
	return d.list.Size()
}

func (d *ListDeque) Empty() bool {
	return d.list.Size() == 0
}
