package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack
	}{
		{
			name:  "should correctly work with SliceStack",
			stack: NewSliceStack(),
		},
		{
			name:  "should correctly work with ListStack",
			stack: NewListStack(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Push(1)
			tt.stack.Push(2)

			assert.False(t, tt.stack.Empty())
			assert.Equal(t, 2, tt.stack.Size())
			assert.Equal(t, 2, tt.stack.Top())
			assert.Equal(t, 2, tt.stack.Pop())
			assert.Equal(t, 1, tt.stack.Size())
			assert.Equal(t, 1, tt.stack.Top())
			assert.Equal(t, 1, tt.stack.Pop())
			assert.Nil(t, tt.stack.Top())
			assert.Nil(t, tt.stack.Pop())
			assert.True(t, tt.stack.Empty())
		})
	}
}
