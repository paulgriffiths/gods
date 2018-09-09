package graphs

import "github.com/paulgriffiths/gods/stacks"

// StackVertex implements a stack of vertex elements.
type StackVertex struct {
	stack stacks.StackInterface
}

// NewStackVertex creates a new stack of vertex elements.
func NewStackVertex() StackVertex {
	return StackVertex{stacks.NewStackInterface()}
}

// Push pushes a new vertex element onto the stack.
func (s *StackVertex) Push(n Vertex) {
	s.stack.Push(n)
}

// Pop pops the top vertex element from the stack.
func (s *StackVertex) Pop() Vertex {
	return s.stack.Pop().(Vertex)
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackVertex) IsEmpty() bool {
	return s.stack.IsEmpty()
}
