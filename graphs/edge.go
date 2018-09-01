package graphs

import (
	"fmt"
)

// Edge represents an undirected graph edge.
type Edge struct {
	V [2]Vertex
}

// NewEdge creates a new undirected edge between the specified vertices.
func NewEdge(v1, v2 Vertex) Edge {
	if v1 == v2 {
		panic("self loops not permitted")
	}
	if v1 <= v2 {
		return Edge{[2]Vertex{v1, v2}}
	}
	return Edge{[2]Vertex{v2, v1}}
}

// Equals returns true if the edge is equal to the provided edge.
func (e Edge) Equals(other Edge) bool {
	return e == other
}

// Less returs true if the edge is less than the provided edge.
func (e Edge) Less(other Edge) bool {
	if e.V[0] < other.V[0] {
		return true
	} else if e.V[0] == other.V[0] && e.V[1] < other.V[1] {
		return true
	}
	return false
}

// String returns a string representation of an undirected edge.
func (e Edge) String() string {
	return fmt.Sprintf("(%d,%d)", e.V[0], e.V[1])
}
