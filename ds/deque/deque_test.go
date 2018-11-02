package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name  string
		deque Deque
	}{
		{
			name:  "should correctly work with SliceDeque",
			deque: NewSliceDeque(),
		},
		{
			name:  "should correctly work with ListDeque",
			deque: NewListDeque(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deque.PushBack(2)
			tt.deque.PushFront(1)
			tt.deque.PushBack(3)
			assert.Equal(t, 3, tt.deque.Back())
			assert.Equal(t, 1, tt.deque.Front())
			assert.Equal(t, 3, tt.deque.PopBack())
			assert.Equal(t, 2, tt.deque.PopBack())
			assert.Equal(t, 1, tt.deque.PopFront())
			assert.Nil(t, tt.deque.PopBack())
			assert.Nil(t, tt.deque.PopFront())
			assert.Nil(t, tt.deque.Back())
			assert.Nil(t, tt.deque.Front())
		})
	}
}
