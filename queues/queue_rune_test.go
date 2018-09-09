package queues_test

import (
	"github.com/paulgriffiths/gods/queues"
	"testing"
)

func TestQueueRuneEmpty(t *testing.T) {
	q := queues.NewQueueRune()
	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}

func TestQueueRuneEnqueueNotEmpty(t *testing.T) {
	q := queues.NewQueueRune()
	q.Enqueue('a')
	if q.IsEmpty() {
		t.Errorf("queue is empty")
	}
}

func TestQueueRuneEnqueueDequeue(t *testing.T) {
	ops := []struct {
		op    string
		value rune
	}{
		{"enqueue", 'a'},
		{"dequeue", 'a'},
		{"enqueue", 'b'},
		{"enqueue", 'c'},
		{"dequeue", 'b'},
		{"dequeue", 'c'},
		{"enqueue", 'd'},
		{"enqueue", 'e'},
		{"enqueue", 'f'},
		{"dequeue", 'd'},
		{"enqueue", 'g'},
		{"dequeue", 'e'},
		{"dequeue", 'f'},
		{"enqueue", 'h'},
		{"dequeue", 'g'},
		{"dequeue", 'h'},
	}

	q := queues.NewQueueRune()
	for _, op := range ops {
		switch op.op {
		case "enqueue":
			q.Enqueue(op.value)
		case "dequeue":
			if c := q.Dequeue(); c != op.value {
				t.Errorf("got %q, want %q", c, op.value)
			}
		}
	}

	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}
