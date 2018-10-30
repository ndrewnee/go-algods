package queue

type SliceQueue struct {
	values []interface{}
}

func New() *SliceQueue {
	return &SliceQueue{}
}

func (q *SliceQueue) Enqueue(value interface{}) {
	q.values = append(q.values, value)
}

func (q *SliceQueue) Dequeue() interface{} {
	if len(q.values) == 0 {
		return nil
	}

	value := q.values[0]
	q.values = q.values[1:]
	return value
}