package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestALGraph(t *testing.T) {
	g := graphs.NewAlGraph(4)
	g.InsertEdge(0, 1)
	g.InsertEdge(0, 2)
	g.InsertEdge(1, 3)
	g.InsertEdge(2, 3)

	a := g.Adjacent(0)
	if !a.Equals(graphs.VertexList{1, 2}) {
		t.Errorf("want %v, got %v", graphs.VertexList{1, 2}, a)
	}

	a = g.Adjacent(1)
	if !a.Equals(graphs.VertexList{0, 3}) {
		t.Errorf("want %v, got %v", graphs.VertexList{0, 3}, a)
	}

	a = g.Adjacent(2)
	if !a.Equals(graphs.VertexList{0, 3}) {
		t.Errorf("want %v, got %v", graphs.VertexList{0, 3}, a)
	}

	a = g.Adjacent(3)
	if !a.Equals(graphs.VertexList{1, 2}) {
		t.Errorf("want %v, got %v", graphs.VertexList{1, 2}, a)
	}
}
