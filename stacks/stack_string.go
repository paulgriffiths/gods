package stacks

// StackString implements a stack of string elements.
type StackString struct {
	stack StackInterface
}

// NewStackString creates a new stack of string elements.
func NewStackString() StackString {
	return StackString{NewStackInterface()}
}

// Push pushes a new string element onto the stack.
func (s *StackString) Push(n string) {
	s.stack.Push(n)
}

// Pop pops the top string element from the stack.
func (s *StackString) Pop() string {
	return s.stack.Pop().(string)
}

// Peek returns the top string element from the stack without
// removing it.
func (s *StackString) Peek() string {
	return s.stack.Peek().(string)
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackString) IsEmpty() bool {
	return s.stack.IsEmpty()
}
