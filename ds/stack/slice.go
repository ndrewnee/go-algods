package stack

type SliceStack struct {
	values []interface{}
}

func NewSliceStack() Stack                   { return &SliceStack{} }
func (s *SliceStack) Size() int              { return len(s.values) }
func (s *SliceStack) IsEmpty() bool          { return len(s.values) == 0 }
func (s *SliceStack) Reset()                 { s.values = nil }
func (s *SliceStack) Push(value interface{}) { s.values = append(s.values, value) }

func (s *SliceStack) Pop() interface{} {
	length := len(s.values)
	if length == 0 {
		return nil
	}

	value := s.values[length-1]
	s.values = s.values[:length-1]
	return value
}
