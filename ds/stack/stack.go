package stack

type Stack interface {
	Push(value interface{})
	Pop() interface{}
	Top() interface{}
	Size() int
	Empty() bool
}
