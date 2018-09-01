package graphs

import (
	"github.com/paulgriffiths/gods/graphs/internal/amGraph"
)

// Graph represents a graph type.
type Graph interface {
	InsertEdge(v1, v2 int)
	Adjacent(v1 int) []int
	Visit(v int)
	IsVisited(v int) bool
	NumVertices() int
}

// NewAMGraph returns a new graph with vertices 0...n-1 and no edges,
// implemented as an adjacency matrix.
func NewAMGraph(n int) Graph {
	return amGraph.NewAmGraph(n)
}
