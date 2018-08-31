package stacks

// StackInt implements a stack of int elements.
type StackInt struct {
	elems []int
	top   int
}

// NewStackInt creates a new stack of int elements.
func NewStackInt() StackInt {
	return StackInt{[]int{}, -1}
}

// Push pushes a new int element onto the stack.
func (s *StackInt) Push(n int) {
	if s.top < len(s.elems) {
		s.elems = append(s.elems, n)
	} else {
		s.elems[s.top+1] = n
	}
	s.top++
}

// Pop pops the top int element from the stack.
func (s *StackInt) Pop() int {
	if s.top < 0 {
		panic("stack underflow")
	}
	n := s.elems[s.top]
	s.top--
	return n
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackInt) IsEmpty() bool {
	return s.top < 0
}
