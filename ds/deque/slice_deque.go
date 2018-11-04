package deque

var _ Deque = &SliceDeque{}

type SliceDeque struct {
	values []interface{}
}

func NewSliceDeque() *SliceDeque {
	return &SliceDeque{}
}

func (d *SliceDeque) PushBack(value interface{}) {
	d.values = append(d.values, value)
}

func (d *SliceDeque) PushFront(value interface{}) {
	d.values = append([]interface{}{value}, d.values...)
}

func (d *SliceDeque) PopBack() interface{} {
	length := len(d.values)
	if length == 0 {
		return nil
	}

	value := d.values[length-1]
	d.values = d.values[:length-1]
	return value
}

func (d *SliceDeque) PopFront() interface{} {
	if len(d.values) == 0 {
		return nil
	}

	value := d.values[0]
	d.values = d.values[1:]
	return value
}

func (d *SliceDeque) Back() interface{} {
	length := len(d.values)
	if length == 0 {
		return nil
	}

	return d.values[length-1]
}

func (d *SliceDeque) Front() interface{} {
	if len(d.values) == 0 {
		return nil
	}

	return d.values[0]
}

func (d *SliceDeque) Size() int {
	return len(d.values)
}

func (d *SliceDeque) Empty() bool {
	return len(d.values) == 0
}
