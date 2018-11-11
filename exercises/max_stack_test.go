package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxStackHandler(t *testing.T) {
	queries := []string{
		"push 2",
		"push 3",
		"push 9",
		"push 7",
		"push 2",
		"max",
		"max",
		"max",
		"pop",
		"max",
	}

	MaxStackHandler(queries)
}

func TestMaxStack(t *testing.T) {
	maxStack := NewMaxStack()

	maxStack.Push(3)
	maxStack.Push(2)
	assert.Equal(t, 3, maxStack.Max())

	maxStack.Push(7)
	maxStack.Push(9)
	assert.Equal(t, 9, maxStack.Max())

	maxStack.Pop()
	assert.Equal(t, 7, maxStack.Max())
}
