package queues

// QueueInterface implements a queue of interface{} elements.
type QueueInterface struct {
	elems []interface{}
}

// NewQueueInterface creates a new queue of interface{} elements.
func NewQueueInterface() QueueInterface {
	return QueueInterface{[]interface{}{}}
}

// Enqueue enqueues an interface{} to a queue.
func (q *QueueInterface) Enqueue(n interface{}) {
	q.elems = append(q.elems, n)
}

// Dequeue dequeues an interface{} from a queue.
func (q *QueueInterface) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	n := q.elems[0]
	q.elems = q.elems[1:len(q.elems)]
	return n
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *QueueInterface) IsEmpty() bool {
	return len(q.elems) == 0
}
