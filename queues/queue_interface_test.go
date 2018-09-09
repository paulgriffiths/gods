package queues_test

import (
	"github.com/paulgriffiths/gods/queues"
	"testing"
)

type mockType struct {
	n int
	s string
}

func (t mockType) Equals(other mockType) bool {
	return t.n == other.n && t.s == other.s
}

func TestQueueInterfaceEmpty(t *testing.T) {
	q := queues.NewQueueInterface()
	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}

func TestQueueInterfaceEnqueueNotEmpty(t *testing.T) {
	q := queues.NewQueueInterface()
	q.Enqueue(mockType{42, "forty two"})
	if q.IsEmpty() {
		t.Errorf("queue is empty")
	}
}

func TestQueueInterfaceEnqueueDequeue(t *testing.T) {
	ops := []struct {
		op    string
		value mockType
	}{
		{"enqueue", mockType{7, "seven"}},
		{"dequeue", mockType{7, "seven"}},
		{"enqueue", mockType{13, "thirteen"}},
		{"enqueue", mockType{18, "eighteen"}},
		{"dequeue", mockType{13, "thirteen"}},
		{"dequeue", mockType{18, "eighteen"}},
		{"enqueue", mockType{11, "eleven"}},
		{"enqueue", mockType{22, "twenty two"}},
		{"enqueue", mockType{33, "thirty three"}},
		{"dequeue", mockType{11, "eleven"}},
		{"enqueue", mockType{44, "forty four"}},
		{"dequeue", mockType{22, "twenty two"}},
		{"dequeue", mockType{33, "thirty three"}},
		{"enqueue", mockType{55, "fifty five"}},
		{"dequeue", mockType{44, "forty four"}},
		{"dequeue", mockType{55, "fifty five"}},
	}

	q := queues.NewQueueInterface()
	for _, op := range ops {
		switch op.op {
		case "enqueue":
			q.Enqueue(op.value)
		case "dequeue":
			if c := q.Dequeue().(mockType); !c.Equals(op.value) {
				t.Errorf("got %v, want %v", c, op.value)
			}
		}
	}

	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}
