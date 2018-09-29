package stacks

// StackFloat implements a stack of float elements.
type StackFloat struct {
	stack StackInterface
}

// NewStackFloat creates a new stack of float elements.
func NewStackFloat() StackFloat {
	return StackFloat{NewStackInterface()}
}

// Push pushes a new float element onto the stack.
func (s *StackFloat) Push(n float64) {
	s.stack.Push(n)
}

// Pop pops the top float element from the stack.
func (s *StackFloat) Pop() float64 {
	return s.stack.Pop().(float64)
}

// Peek returns the top float element from the stack without
// removing it.
func (s *StackFloat) Peek() float64 {
	return s.stack.Peek().(float64)
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackFloat) IsEmpty() bool {
	return s.stack.IsEmpty()
}
