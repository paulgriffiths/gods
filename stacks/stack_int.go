package stacks

// StackInt implements a stack of int elements.
type StackInt struct {
	elems []int
}

// NewStackInt creates a new stack of int elements.
func NewStackInt() StackInt {
	return StackInt{[]int{}}
}

// Push pushes a new int element onto the stack.
func (s *StackInt) Push(n int) {
	s.elems = append(s.elems, n)
}

// Pop pops the top int element from the stack.
func (s *StackInt) Pop() int {
	if len(s.elems) == 0 {
		panic("stack underflow")
	}
	n := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return n
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackInt) IsEmpty() bool {
	return len(s.elems) == 0
}
