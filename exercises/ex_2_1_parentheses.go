package exercises

import (
	"github.com/ndrewnee/go-algods/ds/stack"
)

type Parentheses struct {
	Rune  rune
	Index int
}

func IsValidParentheses(str string) (Parentheses, bool) {
	stack := stack.NewSliceStack()
	for i, r := range str {
		if r == '(' || r == '[' || r == '{' {
			stack.Push(Parentheses{Rune: r, Index: i})
			continue
		}

		if r != ')' && r != ']' && r != '}' {
			continue
		}

		if stack.Empty() {
			return Parentheses{Rune: r, Index: i}, false
		}

		top := (stack.Pop().(Parentheses)).Rune
		if top == '(' && r != ')' || top == '[' && r != ']' || top == '{' && r != '}' {
			return Parentheses{Rune: r, Index: i}, false
		}
	}

	ok := stack.Empty()
	if !ok {
		top := stack.Top().(Parentheses)
		return top, false
	}

	return Parentheses{}, true
}
