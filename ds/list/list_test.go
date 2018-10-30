package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	list := New()

	n1 := list.PushBack(1)
	n2 := list.PushFront(2)
	n3 := list.PushFront(3)
	n4 := list.InsertAfter(n2, 4)
	n5 := list.InsertBefore(n4, 5)
	n6 := list.PushBack(6)
	n7 := list.InsertAfter(n5, 7)

	t.Run("should handle nil nodes", func(t *testing.T) {
		list.Remove(nil)
		assert.Nil(t, list.InsertAfter(nil, -1))
		assert.Nil(t, list.InsertBefore(nil, -2))
	})

	t.Run("should correctly add elements to list", func(t *testing.T) {
		assert.Equal(t, 7, list.Size())
		assert.Equal(t, n3, list.Head())
		assert.Equal(t, n6, list.Tail())
		assert.Equal(t, n7, n5.Next())
		assert.Equal(t, n4, n1.Prev())
	})

	t.Run("should be in right order", func(t *testing.T) {
		values := []int{3, 2, 5, 7, 4, 1, 6}
		node := list.Head()
		for _, value := range values {
			require.NotNil(t, node)
			assert.Equal(t, value, node.Value())
			node = node.Next()
		}
	})

	t.Run("should remove till Node = 4", func(t *testing.T) {
		found := list.Traverse(func(node *Node) bool {
			if node.Value() == 4 {
				return true
			}

			list.Remove(node)
			return false
		}, false)

		assert.Equal(t, 3, list.Size())
		assert.Equal(t, n4, found)
	})

	t.Run("should remove left nodes in reverse", func(t *testing.T) {
		afterRemove := list.Traverse(func(node *Node) bool {
			list.Remove(node)
			return false
		}, true)

		assert.Equal(t, 0, list.Size())
		assert.Nil(t, afterRemove)
	})
}
