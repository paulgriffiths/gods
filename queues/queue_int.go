package queues

// QueueInt implements a queue of int elements.
type QueueInt struct {
	queue QueueInterface
}

// NewQueueInt creates a new queue of int elements.
func NewQueueInt() QueueInt {
	return QueueInt{NewQueueInterface()}
}

// Enqueue enqueues an int to a queue.
func (q *QueueInt) Enqueue(n int) {
	q.queue.Enqueue(n)
}

// Dequeue dequeues an int from a queue.
func (q *QueueInt) Dequeue() int {
	return q.queue.Dequeue().(int)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueInt) IsEmpty() bool {
	return q.queue.IsEmpty()
}
