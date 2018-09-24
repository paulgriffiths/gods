package stacks_test

import (
	"github.com/paulgriffiths/gods/stacks"
	"testing"
)

func TestStackFloatEmpty(t *testing.T) {
	s := stacks.NewStackFloat()
	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}

func TestStackFloatPushNotEmpty(t *testing.T) {
	s := stacks.NewStackFloat()
	s.Push(42.5)
	if s.IsEmpty() {
		t.Errorf("stack is empty")
	}
}

func TestStackFloatPushPop(t *testing.T) {
	ops := []struct {
		op    string
		value float64
	}{
		{"push", 7.1},
		{"pop", 7.1},
		{"push", 13.2},
		{"push", 18.3},
		{"pop", 18.3},
		{"pop", 13.2},
		{"push", 11.1},
		{"push", 22.2},
		{"push", 33.3},
		{"pop", 33.3},
		{"push", 44.4},
		{"pop", 44.4},
		{"pop", 22.2},
		{"push", 55.5},
		{"pop", 55.5},
		{"pop", 11.1},
	}

	s := stacks.NewStackFloat()
	for _, op := range ops {
		switch op.op {
		case "push":
			s.Push(op.value)
		case "pop":
			if c := s.Pop(); c != op.value {
				t.Errorf("got %f, want %f", c, op.value)
			}
		}
	}

	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}
