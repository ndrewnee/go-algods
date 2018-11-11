package exercises

import "github.com/ndrewnee/go-algods/ds/stack"

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

type (
	PairMin struct {
		Value int
		Min   int
	}

	MinQueue struct {
		stack1 stack.Stack
		stack2 stack.Stack
	}
)

func NewMinQueue() *MinQueue {
	return &MinQueue{stack1: stack.NewSliceStack(), stack2: stack.NewSliceStack()}
}

func (mq *MinQueue) Min() int {
	if mq.stack1.Empty() {
		return getMin(mq.stack2.Top())
	}

	if mq.stack2.Empty() {
		return getMin(mq.stack1.Top())
	}

	min1 := getMin(mq.stack1.Top())
	min2 := getMin(mq.stack2.Top())

	return minOf(min1, min2)
}

func (mq *MinQueue) Push(value int) {
	var min int
	if mq.stack1.Empty() {
		min = value
	} else {
		min = minOf(value, getMin(mq.stack1.Top()))
	}

	mq.stack1.Push(PairMin{Value: value, Min: min})
}

//nolint: unparam
func (mq *MinQueue) Pop(value int) {
	defer mq.stack2.Pop()
	if !mq.stack2.Empty() {
		return
	}

	for !mq.stack1.Empty() {
		value := mq.stack1.Pop().(PairMin).Value
		var min int
		if mq.stack2.Empty() {
			min = value
		} else {
			min = minOf(value, getMin(mq.stack2.Top()))
		}
		mq.stack2.Push(PairMin{Value: value, Min: min})
	}
}

func minOf(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMin(value interface{}) int {
	return value.(PairMin).Min
}
