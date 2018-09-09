package stacks_test

import (
	"github.com/paulgriffiths/gods/stacks"
	"testing"
)

type mockType struct {
	n int
	s string
}

func (t mockType) Equals(other mockType) bool {
	return t.n == other.n && t.s == other.s
}

func TestStackInterfaceEmpty(t *testing.T) {
	s := stacks.NewStackInterface()
	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}

func TestStackInterfacePushNotEmpty(t *testing.T) {
	s := stacks.NewStackInterface()
	s.Push(mockType{1, "one"})
	if s.IsEmpty() {
		t.Errorf("stack is empty")
	}
}

func TestStackInterfacePushPop(t *testing.T) {
	ops := []struct {
		op    string
		value mockType
	}{
		{"push", mockType{7, "seven"}},
		{"pop", mockType{7, "seven"}},
		{"push", mockType{13, "thirteen"}},
		{"push", mockType{18, "eighteen"}},
		{"pop", mockType{18, "eighteen"}},
		{"pop", mockType{13, "thirteen"}},
		{"push", mockType{11, "eleven"}},
		{"push", mockType{22, "twenty two"}},
		{"push", mockType{33, "thirty three"}},
		{"pop", mockType{33, "thirty three"}},
		{"push", mockType{44, "forty four"}},
		{"pop", mockType{44, "forty four"}},
		{"pop", mockType{22, "twenty two"}},
		{"push", mockType{55, "fifty five"}},
		{"pop", mockType{55, "fifty five"}},
		{"pop", mockType{11, "eleven"}},
	}

	s := stacks.NewStackInterface()
	for _, op := range ops {
		switch op.op {
		case "push":
			s.Push(op.value)
		case "pop":
			c := s.Pop().(mockType)
			if !c.Equals(op.value) {
				t.Errorf("got %v, want %v", c, op.value)
			}
		}
	}

	if !s.IsEmpty() {
		t.Errorf("stack is not empty")
	}
}
