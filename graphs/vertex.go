package graphs

import "fmt"

// Vertex represents a graph vertex.
type Vertex int

// Equals returns true if the vertex is equal to the provided vertex.
func (v Vertex) Equals(other Vertex) bool {
	return v == other
}

// Less returns true if the vertex is less than the provided vertex.
func (v Vertex) Less(other Vertex) bool {
	return v < other
}

// String returns a string representation of a vertex.
func (v Vertex) String() string {
	return fmt.Sprintf("{%d}", v)
}
