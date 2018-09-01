package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestBfs(t *testing.T) {
	g := graphs.NewAmGraph(8)
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
	l := graphs.BfsVertexList(g, 0)
	expected := graphs.VertexList{0, 1, 2, 3, 4, 5, 6, 7}
	if !l.Equals(expected) {
		t.Errorf("got %v, want %v", l, expected)
	}
}

func TestDfs(t *testing.T) {
	g := graphs.NewAmGraph(8)
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
	l := graphs.DfsVertexList(g, 0)
	expected := graphs.VertexList{0, 1, 3, 7, 4, 5, 2, 6}
	if !l.Equals(expected) {
		t.Errorf("got %v, want %v", l, expected)
	}
}

func TestDfsIter(t *testing.T) {
	g := graphs.NewAmGraph(8)
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
	l := graphs.DfsVertexListIterative(g, 0)
	expected := graphs.VertexList{0, 1, 3, 7, 4, 5, 2, 6}
	if !l.Equals(expected) {
		t.Errorf("got %v, want %v", l, expected)
	}
}
