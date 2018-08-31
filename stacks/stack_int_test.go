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

func TestPopEmptyStackInt(t *testing.T) {
	s := stacks.NewStackInt()
	s.Push(42)
	s.Pop()
	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}

func TestPushPopStackInt(t *testing.T) {
	s := stacks.NewStackInt()
	vals := []int{
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	}

	for _, n := range vals {
		s.Push(n)
	}

	for n := len(vals) - 1; n >= 0; n-- {
		p := s.Pop()
		if p != vals[n] {
			t.Errorf("got %d, want %d", p, vals[n])
		}
	}
}
