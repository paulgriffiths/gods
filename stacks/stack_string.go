package stacks

// StackString implements a stack of string elements.
type StackString struct {
	elems []string
}

// NewStackString creates a new stack of string elements.
func NewStackString() StackString {
	return StackString{[]string{}}
}

// Push pushes a new string element onto the stack.
func (s *StackString) Push(n string) {
	s.elems = append(s.elems, n)
}

// Pop pops the top string element from the stack.
func (s *StackString) Pop() string {
	if len(s.elems) == 0 {
		panic("stack underflow")
	}
	n := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return n
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackString) IsEmpty() bool {
	return len(s.elems) == 0
}
