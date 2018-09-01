package queues

// QueueInt implements a queue of int elements.
type QueueInt struct {
	elems []int
}

// NewQueueInt creates a new queue of int elements.
func NewQueueInt() QueueInt {
	return QueueInt{[]int{}}
}

// Enqueue enqueues an int to a queue.
func (q *QueueInt) Enqueue(n int) {
	q.elems = append(q.elems, n)
}

// Dequeue dequeues an int from a queue.
func (q *QueueInt) Dequeue() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	n := q.elems[0]
	q.elems = q.elems[1:len(q.elems)]
	return n
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueInt) IsEmpty() bool {
	return len(q.elems) == 0
}
