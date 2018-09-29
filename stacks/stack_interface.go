package stacks

// StackInterface implements a stack of interface{} elements.
type StackInterface struct {
	elems []interface{}
}

// NewStackInterface creates a new stack of interface{} elements.
func NewStackInterface() StackInterface {
	return StackInterface{[]interface{}{}}
}

// Push pushes a new interface{} element onto the stack.
func (s *StackInterface) Push(n interface{}) {
	s.elems = append(s.elems, n)
}

// Pop pops the top interface{} element from the stack.
func (s *StackInterface) Pop() interface{} {
	if len(s.elems) == 0 {
		panic("stack underflow")
	}
	n := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return n
}

// Peek returns the top interface{} element from the stack without
// removing it.
func (s *StackInterface) Peek() interface{} {
	if len(s.elems) == 0 {
		panic("stack underflow")
	}
	n := s.elems[len(s.elems)-1]
	return n
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackInterface) IsEmpty() bool {
	return len(s.elems) == 0
}
