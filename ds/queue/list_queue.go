package queue

import "github.com/ndrewnee/go-algods/ds/list"

var _ Queue = &ListQueue{}

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

func (q *ListQueue) Size() int {
	return q.list.Size()
}

func (q *ListQueue) Empty() bool {
	return q.list.Size() == 0
}
