package stack

import "github.com/ndrewnee/go-algods/ds/list"

type ListStack struct {
	list *list.List
}

func NewListStack() Stack                   { return &ListStack{list: list.New()} }
func (s *ListStack) Size() int              { return s.list.Size() }
func (s *ListStack) IsEmpty() bool          { return s.list.Size() == 0 }
func (s *ListStack) Reset()                 { s.list = list.New() }
func (s *ListStack) Push(value interface{}) { _ = s.list.PushFront(value) }

func (s *ListStack) Pop() interface{} {
	if head := s.list.Head(); head != nil {
		return head.Value()
	}
	return nil
}
