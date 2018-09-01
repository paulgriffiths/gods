package graphs

// QueueVertex implements a queue of Vertex elements.
type QueueVertex struct {
	elems []Vertex
}

// NewQueueVertex creates a new queue of Vertex elements.
func NewQueueVertex() QueueVertex {
	return QueueVertex{[]Vertex{}}
}

// Enqueue enqueues an Vertex to a queue.
func (q *QueueVertex) Enqueue(n Vertex) {
	q.elems = append(q.elems, n)
}

// Dequeue dequeues an Vertex from a queue.
func (q *QueueVertex) Dequeue() Vertex {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	n := q.elems[0]
	q.elems = q.elems[1:len(q.elems)]
	return n
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueVertex) IsEmpty() bool {
	return len(q.elems) == 0
}
