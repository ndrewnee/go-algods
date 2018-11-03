package exercises

import (
	"github.com/ndrewnee/go-algods/ds/stack"
)

func IsValidParentheses(str string) bool {
	stack := stack.NewSliceStack()
	for _, r := range str {
		if r == '(' || r == '[' || r == '{' {
			stack.Push(r)
			continue
		}

		if stack.Size() == 0 {
			return false
		}

		topRaw := stack.Pop()
		top := topRaw.(rune)
		if top == '(' && r != ')' || top == '[' && r != ']' || top == '{' && r != '}' {
			return false
		}
	}

	return stack.Size() == 0
}
