package stacks_test

import (
	"github.com/paulgriffiths/gods/stacks"
	"testing"
)

func TestStackRuneEmpty(t *testing.T) {
	s := stacks.NewStackRune()
	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}

func TestStackRunePushNotEmpty(t *testing.T) {
	s := stacks.NewStackRune()
	s.Push('a')
	if s.IsEmpty() {
		t.Errorf("stack is empty")
	}
}

func TestStackRunePushPop(t *testing.T) {
	ops := []struct {
		op    string
		value rune
	}{
		{"push", 'a'},
		{"pop", 'a'},
		{"push", 'b'},
		{"push", 'c'},
		{"pop", 'c'},
		{"pop", 'b'},
		{"push", 'd'},
		{"push", 'e'},
		{"push", 'f'},
		{"pop", 'f'},
		{"push", 'g'},
		{"pop", 'g'},
		{"pop", 'e'},
		{"push", 'h'},
		{"pop", 'h'},
		{"pop", 'd'},
	}

	s := stacks.NewStackRune()
	for _, op := range ops {
		switch op.op {
		case "push":
			s.Push(op.value)
		case "pop":
			if c := s.Pop(); c != op.value {
				t.Errorf("got %q, want %q", c, op.value)
			}
		}
	}

	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}
