package graphs

// StackVertex implements a stack of Vertex elements.
type StackVertex struct {
	elems []Vertex
}

// NewStackVertex creates a new stack of Vertex elements.
func NewStackVertex() StackVertex {
	return StackVertex{[]Vertex{}}
}

// Push pushes a new Vertex element onto the stack.
func (s *StackVertex) Push(n Vertex) {
	s.elems = append(s.elems, n)
}

// Pop pops the top Vertex element from the stack.
func (s *StackVertex) Pop() Vertex {
	if len(s.elems) == 0 {
		panic("stack underflow")
	}
	n := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return n
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackVertex) IsEmpty() bool {
	return len(s.elems) == 0
}
