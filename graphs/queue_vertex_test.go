package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestEmptyQueueVertex(t *testing.T) {
	q := graphs.NewQueueVertex()
	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}

func TestEnqueueNotEmptyQueueVertex(t *testing.T) {
	q := graphs.NewQueueVertex()
	q.Enqueue(42)
	if q.IsEmpty() {
		t.Errorf("queue is empty")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	ops := []struct {
		op    string
		value graphs.Vertex
	}{
		{"enqueue", 7},
		{"dequeue", 7},
		{"enqueue", 13},
		{"enqueue", 18},
		{"dequeue", 13},
		{"dequeue", 18},
		{"enqueue", 11},
		{"enqueue", 22},
		{"enqueue", 33},
		{"dequeue", 11},
		{"enqueue", 44},
		{"dequeue", 22},
		{"dequeue", 33},
		{"enqueue", 55},
		{"dequeue", 44},
		{"dequeue", 55},
	}

	q := graphs.NewQueueVertex()
	for _, op := range ops {
		switch op.op {
		case "enqueue":
			q.Enqueue(op.value)
		case "dequeue":
			c := q.Dequeue()
			if c != op.value {
				t.Errorf("got %d, want %d", c, op.value)
			}
		}
	}

	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}
