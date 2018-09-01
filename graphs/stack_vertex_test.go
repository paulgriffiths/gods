package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestEmptyStackVertex(t *testing.T) {
	s := graphs.NewStackVertex()
	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}

func TestPushNotEmptyStackVertex(t *testing.T) {
	s := graphs.NewStackVertex()
	s.Push(42)
	if s.IsEmpty() {
		t.Errorf("stack is empty")
	}
}

func TestPushPop(t *testing.T) {
	ops := []struct {
		op    string
		value graphs.Vertex
	}{
		{"push", 7},
		{"pop", 7},
		{"push", 13},
		{"push", 18},
		{"pop", 18},
		{"pop", 13},
		{"push", 11},
		{"push", 22},
		{"push", 33},
		{"pop", 33},
		{"push", 44},
		{"pop", 44},
		{"pop", 22},
		{"push", 55},
		{"pop", 55},
		{"pop", 11},
	}

	s := graphs.NewStackVertex()
	for _, op := range ops {
		switch op.op {
		case "push":
			s.Push(op.value)
		case "pop":
			c := s.Pop()
			if c != op.value {
				t.Errorf("got %d, want %d", c, op.value)
			}
		}
	}

	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}
