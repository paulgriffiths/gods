package stacks

// StackRune implements a stack of rune elements.
type StackRune struct {
	elems []rune
}

// NewStackRune creates a new stack of rune elements.
func NewStackRune() StackRune {
	return StackRune{[]rune{}}
}

// Push pushes a new rune element onto the stack.
func (s *StackRune) Push(n rune) {
	s.elems = append(s.elems, n)
}

// Pop pops the top rune element from the stack.
func (s *StackRune) Pop() rune {
	if len(s.elems) == 0 {
		panic("stack underflow")
	}
	n := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return n
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackRune) IsEmpty() bool {
	return len(s.elems) == 0
}
