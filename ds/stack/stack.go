package stack

type Stack interface {
	Size() int
	IsEmpty() bool
	Reset()
	Push(value interface{})
	Pop() interface{}
}
