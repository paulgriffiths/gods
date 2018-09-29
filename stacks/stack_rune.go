package stacks

// StackRune implements a stack of rune elements.
type StackRune struct {
	stack StackInterface
}

// NewStackRune creates a new stack of rune elements.
func NewStackRune() StackRune {
	return StackRune{NewStackInterface()}
}

// Push pushes a new rune element onto the stack.
func (s *StackRune) Push(n rune) {
	s.stack.Push(n)
}

// Pop pops the top rune element from the stack.
func (s *StackRune) Pop() rune {
	return s.stack.Pop().(rune)
}

// Peek returns the top rune element from the stack without
// removing it.
func (s *StackRune) Peek() rune {
	return s.stack.Peek().(rune)
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackRune) IsEmpty() bool {
	return s.stack.IsEmpty()
}
