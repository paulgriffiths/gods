package stacks

// StackInt implements a stack of int elements.
type StackInt struct {
	stack StackInterface
}

// NewStackInt creates a new stack of int elements.
func NewStackInt() StackInt {
	return StackInt{NewStackInterface()}
}

// Push pushes a new int element onto the stack.
func (s *StackInt) Push(n int) {
	s.stack.Push(n)
}

// Pop pops the top int element from the stack.
func (s *StackInt) Pop() int {
	return s.stack.Pop().(int)
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackInt) IsEmpty() bool {
	return s.stack.IsEmpty()
}
