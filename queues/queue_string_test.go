package queues_test

import (
	"github.com/paulgriffiths/gods/queues"
	"testing"
)

func TestQueueStringEmpty(t *testing.T) {
	q := queues.NewQueueString()
	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}

func TestQueueStringEnqueueNotEmpty(t *testing.T) {
	q := queues.NewQueueString()
	q.Enqueue("spitfire")
	if q.IsEmpty() {
		t.Errorf("queue is empty")
	}
}

func TestQueueStringEnqueueDequeue(t *testing.T) {
	ops := []struct {
		op    string
		value string
	}{
		{"enqueue", "tornado"},
		{"dequeue", "tornado"},
		{"enqueue", "hunter"},
		{"enqueue", "meteor"},
		{"dequeue", "hunter"},
		{"dequeue", "meteor"},
		{"enqueue", "lightning"},
		{"enqueue", "buccaneer"},
		{"enqueue", "canberra"},
		{"dequeue", "lightning"},
		{"enqueue", "harrier"},
		{"dequeue", "buccaneer"},
		{"dequeue", "canberra"},
		{"enqueue", "hawk"},
		{"dequeue", "harrier"},
		{"dequeue", "hawk"},
	}

	q := queues.NewQueueString()
	for _, op := range ops {
		switch op.op {
		case "enqueue":
			q.Enqueue(op.value)
		case "dequeue":
			if c := q.Dequeue(); c != op.value {
				t.Errorf("got %s, want %s", c, op.value)
			}
		}
	}

	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}
