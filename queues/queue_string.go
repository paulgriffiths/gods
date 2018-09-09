package queues

// QueueString implements a queue of string elements.
type QueueString struct {
	queue QueueInterface
}

// NewQueueString creates a new queue of string elements.
func NewQueueString() QueueString {
	return QueueString{NewQueueInterface()}
}

// Enqueue enqueues an string to a queue.
func (q *QueueString) Enqueue(n string) {
	q.queue.Enqueue(n)
}

// Dequeue dequeues an string from a queue.
func (q *QueueString) Dequeue() string {
	return q.queue.Dequeue().(string)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueString) IsEmpty() bool {
	return q.queue.IsEmpty()
}
