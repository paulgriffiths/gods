package stacks_test

import (
	"github.com/paulgriffiths/gods/stacks"
	"testing"
)

func TestStackStringEmpty(t *testing.T) {
	s := stacks.NewStackString()
	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}

func TestStackStringPushNotEmpty(t *testing.T) {
	s := stacks.NewStackString()
	s.Push("battleship")
	if s.IsEmpty() {
		t.Errorf("stack is empty")
	}
}

func TestStackStringPushPop(t *testing.T) {
	ops := []struct {
		op    string
		value string
	}{
		{"push", "destroyer"},
		{"pop", "destroyer"},
		{"push", "cruiser"},
		{"push", "corvette"},
		{"pop", "corvette"},
		{"pop", "cruiser"},
		{"push", "submarine"},
		{"push", "aircraft carrier"},
		{"push", "frigate"},
		{"pop", "frigate"},
		{"push", "submarine"},
		{"pop", "submarine"},
		{"pop", "aircraft carrier"},
		{"push", "battlecruiser"},
		{"pop", "battlecruiser"},
		{"pop", "submarine"},
	}

	s := stacks.NewStackString()
	for _, op := range ops {
		switch op.op {
		case "push":
			s.Push(op.value)
		case "pop":
			if c := s.Pop(); c != op.value {
				t.Errorf("got %s, want %s", c, op.value)
			}
		}
	}

	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}
