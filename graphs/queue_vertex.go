package graphs

import "github.com/paulgriffiths/gods/queues"

// QueueVertex implements a queue of vertex elements.
type QueueVertex struct {
	queue queues.QueueInterface
}

// NewQueueVertex creates a new queue of vertex elements.
func NewQueueVertex() QueueVertex {
	return QueueVertex{queues.NewQueueInterface()}
}

// Enqueue enqueues an vertex to a queue.
func (q *QueueVertex) Enqueue(n Vertex) {
	q.queue.Enqueue(n)
}

// Dequeue dequeues an vertex from a queue.
func (q *QueueVertex) Dequeue() Vertex {
	return q.queue.Dequeue().(Vertex)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueVertex) IsEmpty() bool {
	return q.queue.IsEmpty()
}
