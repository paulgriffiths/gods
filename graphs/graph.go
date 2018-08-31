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

// DfsSpan returns a depth first spanning tree for a graph.
func DfsSpan(g Graph, v int) []int {
	l := []int{}
	f := func(n int) {
		l = append(l, n)
	}
	Dfs(g, v, f)
	return l
}

// Dfs performs a depth first search of a graph.
func Dfs(g Graph, v int, f func(int)) {
	g.Visit(v)
	f(v)
	for _, x := range g.Adjacent(v) {
		if !g.IsVisited(x) {
			Dfs(g, x, f)
		}
	}
}
