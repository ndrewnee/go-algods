package queue

type Queue interface {
	Enqueue(value interface{})
	Dequeue() interface{}
}
