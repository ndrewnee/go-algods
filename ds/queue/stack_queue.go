package queue

import "github.com/ndrewnee/go-algods/ds/stack"

type StackQueue struct {
	left  stack.Stack
	right stack.Stack
}

func NewStackQueue() *StackQueue {
	return &StackQueue{left: stack.NewSliceStack(), right: stack.NewSliceStack()}
}

func (q *StackQueue) Enqueue(value interface{}) {
	q.left.Push(value)
}

func (q *StackQueue) Dequeue() interface{} {
	if q.right.Size() > 0 {
		return q.right.Pop()
	}

	for v := q.left.Pop(); v != nil; v = q.left.Pop() {
		q.right.Push(v)
	}

	return q.right.Pop()
}

func (q *StackQueue) Size() int {
	return q.left.Size() + q.right.Size()
}
