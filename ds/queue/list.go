package queue

import "github.com/ndrewnee/go-algods/ds/list"

type ListQueue struct {
	list *list.List
}

func NewListQueue() *ListQueue {
	return &ListQueue{list: list.New()}
}

func (q *ListQueue) Enqueue(value interface{}) {
	_ = q.list.PushBack(value)
}

func (q *ListQueue) Dequeue() interface{} {
	head := q.list.Head()
	if head == nil {
		return nil
	}

	q.list.Remove(head)
	return head.Value()
}
