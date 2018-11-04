package stack

var _ Stack = &SliceStack{}

type SliceStack struct {
	values []interface{}
}

func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func (s *SliceStack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *SliceStack) Pop() interface{} {
	length := len(s.values)
	if length == 0 {
		return nil
	}

	value := s.values[length-1]
	s.values = s.values[:length-1]
	return value
}

func (s *SliceStack) Top() interface{} {
	length := len(s.values)
	if length == 0 {
		return nil
	}

	value := s.values[length-1]
	return value
}

func (s *SliceStack) Size() int {
	return len(s.values)
}

func (s *SliceStack) Empty() bool {
	return len(s.values) == 0
}
