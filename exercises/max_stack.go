package exercises

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ndrewnee/go-algods/ds/stack"
)

func MaxStackHandler(queries []string) {
	maxStack := NewMaxStack()
	for _, query := range queries {
		if strings.HasPrefix(query, "push ") {
			valueString := strings.TrimPrefix(query, "push ")
			value, err := strconv.Atoi(valueString)
			if err != nil {
				log.Fatal("convert string to int failed", err)
				continue
			}

			maxStack.Push(value)
			continue
		}

		if strings.HasPrefix(query, "pop") {
			maxStack.Pop()
			continue
		}

		if strings.HasPrefix(query, "max") {
			fmt.Println(maxStack.Max())
			continue
		}
	}
}

type MaxStack struct {
	stack stack.Stack
}

func NewMaxStack() *MaxStack {
	return &MaxStack{stack: stack.NewSliceStack()}
}

func (ms *MaxStack) Max() int {
	return getMax(ms.stack.Top())
}

func (ms *MaxStack) Push(value int) {
	var max int
	if ms.stack.Empty() {
		max = value
	} else {
		max = maxOf(value, getMax(ms.stack.Top()))
	}

	ms.stack.Push(PairMax{Value: value, Max: max})
}

func (ms *MaxStack) Pop() {
	ms.stack.Pop()
}
