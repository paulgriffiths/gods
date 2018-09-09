package queues

// QueueRune implements a queue of rune elements.
type QueueRune struct {
	queue QueueInterface
}

// NewQueueRune creates a new queue of rune elements.
func NewQueueRune() QueueRune {
	return QueueRune{NewQueueInterface()}
}

// Enqueue enqueues an rune to a queue.
func (q *QueueRune) Enqueue(n rune) {
	q.queue.Enqueue(n)
}

// Dequeue dequeues an rune from a queue.
func (q *QueueRune) Dequeue() rune {
	return q.queue.Dequeue().(rune)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueRune) IsEmpty() bool {
	return q.queue.IsEmpty()
}
