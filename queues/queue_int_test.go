package queues_test

import (
	"github.com/paulgriffiths/gods/queues"
	"testing"
)

func TestEmptyQueueInt(t *testing.T) {
	q := queues.NewQueueInt()
	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}

func TestEnqueueNotEmptyQueueInt(t *testing.T) {
	q := queues.NewQueueInt()
	q.Enqueue(42)
	if q.IsEmpty() {
		t.Errorf("queue is empty")
	}
}

func TestDequeueEmptyQueueInt(t *testing.T) {
	q := queues.NewQueueInt()
	q.Enqueue(42)
	q.Dequeue()
	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}

func TestDequeueQueueInt(t *testing.T) {
	vals := []int{
		10, 71, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 71, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 71, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 71, 42, 43, 44, 45, 46, 47, 48, 49,
	}
	q := queues.NewQueueInt()
	for _, v := range vals {
		q.Enqueue(v)
	}
	for n, _ := range vals {
		p := q.Dequeue()
		if p != vals[n] {
			t.Errorf("got %d, want %d", p, vals[n])
		}
	}
	if !q.IsEmpty() {
		t.Errorf("queue is not empty")
	}
}
