package exercises

import "github.com/ndrewnee/go-algods/ds/stack"

func FindMinWindows2(numbers []int, windowSize int) []int {
	return findMinWindows(NewMinQueue2(), numbers, windowSize)
}

type (
	Pair struct {
		Value int
		Min   int
	}

	MinQueue2 struct {
		stack1 stack.Stack
		stack2 stack.Stack
	}
)

func NewMinQueue2() *MinQueue2 {
	return &MinQueue2{stack1: stack.NewSliceStack(), stack2: stack.NewSliceStack()}
}

func (mq *MinQueue2) Min() int {
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

func (mq *MinQueue2) Push(value int) {
	var min int
	if mq.stack1.Empty() {
		min = value
	} else {
		min = minOf(value, getMin(mq.stack1.Top()))
	}

	mq.stack1.Push(Pair{Value: value, Min: min})
}

//nolint: unparam
func (mq *MinQueue2) Pop(value int) {
	defer mq.stack2.Pop()
	if !mq.stack2.Empty() {
		return
	}

	for !mq.stack1.Empty() {
		value := mq.stack1.Pop().(Pair).Value
		var min int
		if mq.stack2.Empty() {
			min = value
		} else {
			min = minOf(value, getMin(mq.stack2.Top()))
		}
		mq.stack2.Push(Pair{Value: value, Min: min})
	}
}

func minOf(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMin(value interface{}) int {
	return value.(Pair).Min
}
