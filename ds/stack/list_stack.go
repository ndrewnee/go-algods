package stack

import "github.com/ndrewnee/go-algods/ds/list"

var _ Stack = &ListStack{}

type ListStack struct {
	list *list.List
}

func NewListStack() *ListStack {
	return &ListStack{list: list.New()}
}

func (s *ListStack) Push(value interface{}) {
	_ = s.list.PushFront(value)
}

func (s *ListStack) Pop() interface{} {
	head := s.list.Head()
	if head == nil {
		return nil
	}

	s.list.Remove(head)
	return head.Value()
}
