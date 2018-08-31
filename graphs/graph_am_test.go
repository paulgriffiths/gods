package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func IntSliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for n := 0; n < len(s1); n++ {
		if s1[n] != s2[n] {
			return false
		}
	}

	return true
}

func TestGraph(t *testing.T) {
	g := graphs.NewAMGraph(4)
	if g.NumVertices() != 4 {
		t.Errorf("got %d, want %d", g.NumVertices(), 4)
	}
}

func TestDfs(t *testing.T) {
	g := graphs.NewAMGraph(8)
	g.InsertEdge(0, 1)
	g.InsertEdge(0, 2)
	g.InsertEdge(1, 3)
	g.InsertEdge(1, 3)
	g.InsertEdge(2, 5)
	g.InsertEdge(2, 6)
	g.InsertEdge(3, 7)
	g.InsertEdge(4, 7)
	g.InsertEdge(5, 7)
	g.InsertEdge(6, 7)
	l := graphs.DfsSpan(g, 0)
	if !IntSliceEqual(l, []int{0, 1, 3, 7, 4, 5, 2, 6}) {
		t.Errorf("got %v, want %v", l, []int{0, 1, 3, 7, 4, 5, 2, 6})
	}
}
