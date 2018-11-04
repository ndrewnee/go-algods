package exercises

import "github.com/ndrewnee/go-algods/ds/deque"

func FindMinWindows1(numbers []int, windowSize int) []int {
	return findMinWindows(NewMinQueue1(), numbers, windowSize)
}

type MinQueue1 struct {
	deque deque.Deque
}

func NewMinQueue1() *MinQueue1 {
	return &MinQueue1{deque: deque.NewSliceDeque()}
}

func (md *MinQueue1) Min() int {
	return md.deque.Front().(int)
}

func (md *MinQueue1) Push(value int) {
	for !md.deque.Empty() && md.deque.Back().(int) > value {
		md.deque.PopBack()
	}

	md.deque.PushBack(value)
}

func (md *MinQueue1) Pop(value int) {
	if !md.deque.Empty() && md.deque.Front().(int) == value {
		md.deque.PopFront()
	}
}
