package exercises

import "github.com/ndrewnee/go-algods/ds/deque"

type MinQueue struct {
	deque deque.Deque
}

func NewMinQueue() *MinQueue {
	return &MinQueue{deque: deque.NewSliceDeque()}
}

func (md *MinQueue) Min() int {
	return md.deque.Front().(int)
}

func (md *MinQueue) Push(addedValue int) {
	for !md.deque.Empty() && md.deque.Back().(int) > addedValue {
		md.deque.PopBack()
	}

	md.deque.PushBack(addedValue)
}

func (md *MinQueue) Pop(removedValue int) {
	if !md.deque.Empty() && md.deque.Front().(int) == removedValue {
		md.deque.PopFront()
	}
}

func FindMinWindows(numbers []int, windowSize int) []int {
	length := len(numbers)
	if length == 0 {
		return nil
	}

	if windowSize == 1 {
		return numbers
	}

	if windowSize < 1 || windowSize >= length {
		return []int{findMin(numbers)}
	}

	minQueue := NewMinQueue()
	for i := 0; i < windowSize; i++ {
		minQueue.Push(numbers[i])
	}
	minWindows := []int{minQueue.Min()}

	for i, j := 0, windowSize; j < length; i, j = i+1, j+1 {
		minQueue.Push(numbers[j])
		minQueue.Pop(numbers[i])
		minWindows = append(minWindows, minQueue.Min())
	}

	return minWindows
}

func findMin(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}

	return min
}
