package stacks_test

import (
	"github.com/paulgriffiths/gods/stacks"
	"testing"
)

func TestEmptyStackInt(t *testing.T) {
	s := stacks.NewStackInt()
	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}

func TestPushNotEmptyStackInt(t *testing.T) {
	s := stacks.NewStackInt()
	s.Push(42)
	if s.IsEmpty() {
		t.Errorf("stack is empty")
	}
}

func TestPushPop(t *testing.T) {
	ops := []struct {
		op    string
		value int
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

	s := stacks.NewStackInt()
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
