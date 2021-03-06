package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name  string
		queue Queue
	}{
		{
			name:  "should correctly work with SliceQueue",
			queue: NewSliceQueue(),
		},
		{
			name:  "should correctly work with ListQueue",
			queue: NewListQueue(),
		},
		{
			name:  "should correctly work with StackQueue",
			queue: NewStackQueue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.queue.Enqueue(1)
			tt.queue.Enqueue(2)

			assert.False(t, tt.queue.Empty())
			assert.Equal(t, 2, tt.queue.Size())
			assert.Equal(t, 1, tt.queue.Dequeue())
			assert.Equal(t, 1, tt.queue.Size())
			assert.Equal(t, 2, tt.queue.Dequeue())
			assert.Nil(t, tt.queue.Dequeue())
			assert.True(t, tt.queue.Empty())
		})
	}
}
