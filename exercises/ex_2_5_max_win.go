package exercises

import "github.com/ndrewnee/go-algods/ds/stack"

func FindMaxWindows(numbers []int, windowSize int) []int {
	length := len(numbers)
	if length == 0 {
		return nil
	}

	if windowSize == 1 {
		return numbers
	}

	if windowSize < 1 || windowSize >= length {
		return []int{findMax(numbers)}
	}

	maxQueue := NewMaxQueue()
	for i := 0; i < windowSize; i++ {
		maxQueue.Push(numbers[i])
	}
	maxWindows := []int{maxQueue.Max()}

	for i, j := 0, windowSize; j < length; i, j = i+1, j+1 {
		maxQueue.Push(numbers[j])
		maxQueue.Pop()
		maxWindows = append(maxWindows, maxQueue.Max())
	}

	return maxWindows
}

func findMax(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	return max
}

type (
	PairMax struct {
		Value int
		Max   int
	}

	MaxQueue struct {
		stack1 stack.Stack
		stack2 stack.Stack
	}
)

func NewMaxQueue() *MaxQueue {
	return &MaxQueue{stack1: stack.NewSliceStack(), stack2: stack.NewSliceStack()}
}

func (mq *MaxQueue) Max() int {
	if mq.stack1.Empty() {
		return getMax(mq.stack2.Top())
	}

	if mq.stack2.Empty() {
		return getMax(mq.stack1.Top())
	}

	max1 := getMax(mq.stack1.Top())
	max2 := getMax(mq.stack2.Top())

	return maxOf(max1, max2)
}

func (mq *MaxQueue) Push(value int) {
	var max int
	if mq.stack1.Empty() {
		max = value
	} else {
		max = maxOf(value, getMax(mq.stack1.Top()))
	}

	mq.stack1.Push(PairMax{Value: value, Max: max})
}

func (mq *MaxQueue) Pop() {
	defer mq.stack2.Pop()
	if !mq.stack2.Empty() {
		return
	}

	for !mq.stack1.Empty() {
		value := mq.stack1.Pop().(PairMax).Value
		var max int
		if mq.stack2.Empty() {
			max = value
		} else {
			max = maxOf(value, getMax(mq.stack2.Top()))
		}
		mq.stack2.Push(PairMax{Value: value, Max: max})
	}
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getMax(value interface{}) int {
	return value.(PairMax).Max
}
