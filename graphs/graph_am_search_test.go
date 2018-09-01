package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestBfs(t *testing.T) {
	g := graphs.NewAMGraph(8)
	g.InsertEdge(0, 1)
	g.InsertEdge(0, 2)
	g.InsertEdge(1, 3)
	g.InsertEdge(1, 4)
	g.InsertEdge(2, 5)
	g.InsertEdge(2, 6)
	g.InsertEdge(3, 7)
	g.InsertEdge(4, 7)
	g.InsertEdge(5, 7)
	g.InsertEdge(6, 7)
	l := graphs.BfsSpan(g, 0)
	if !IntSliceEqual(l, []int{0, 1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("got %v, want %v", l, []int{0, 1, 2, 3, 4, 5, 6, 7})
	}
}

func TestDfs(t *testing.T) {
	g := graphs.NewAMGraph(8)
	g.InsertEdge(0, 1)
	g.InsertEdge(0, 2)
	g.InsertEdge(1, 3)
	g.InsertEdge(1, 4)
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

func TestDfsIter(t *testing.T) {
	g := graphs.NewAMGraph(8)
	g.InsertEdge(0, 1)
	g.InsertEdge(0, 2)
	g.InsertEdge(1, 3)
	g.InsertEdge(1, 4)
	g.InsertEdge(2, 5)
	g.InsertEdge(2, 6)
	g.InsertEdge(3, 7)
	g.InsertEdge(4, 7)
	g.InsertEdge(5, 7)
	g.InsertEdge(6, 7)
	l := graphs.DfsSpanIterative(g, 0)
	if !IntSliceEqual(l, []int{0, 1, 3, 7, 4, 5, 2, 6}) {
		t.Errorf("got %v, want %v", l, []int{0, 1, 3, 7, 4, 5, 2, 6})
	}
}
