package queues

import (
	"github.com/paulgriffiths/gods/nodes"
)

// QueueInt implements a queue of int elements.
type QueueInt struct {
	front *nodes.NodeInt
	back  *nodes.NodeInt
}

// NewQueueInt creates a new queue of int elements.
func NewQueueInt() QueueInt {
	return QueueInt{nil, nil}
}

// Enqueue enqueues an int to a queue.
func (q *QueueInt) Enqueue(n int) {
	newNode := &nodes.NodeInt{n, nil}
	if q.front == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.Next = newNode
		q.back = newNode
	}
}

// Dequeue dequeues an int from a queue.
func (q *QueueInt) Dequeue() int {
	if q.front == nil {
		panic("queue is empty")
	}
	n := q.front.Value
	q.front = q.front.Next
	if q.front == nil {
		q.back = nil
	}
	return n
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueInt) IsEmpty() bool {
	return q.front == nil && q.back == nil
}
