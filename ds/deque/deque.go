package deque

type Deque interface {
	PushBack(value interface{})
	PushFront(value interface{})
	PopBack() interface{}
	PopFront() interface{}
	Back() interface{}
	Front() interface{}
	Size() int
	Empty() bool
}
