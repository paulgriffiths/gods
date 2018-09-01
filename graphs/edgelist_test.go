package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestEdgeListLength(t *testing.T) {
	e := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(0, 2),
		graphs.NewEdge(1, 2),
	}
	if e.Len() != 3 {
		t.Errorf("got %d, want %d", e.Len(), 3)
	}
}

func TestEdgeListEqual(t *testing.T) {
	e1 := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(0, 2),
		graphs.NewEdge(1, 2),
	}
	e2 := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(0, 2),
		graphs.NewEdge(1, 2),
	}
	if !e1.Equals(e2) {
		t.Errorf("edge lists don't compare equal")
	}
}

func TestEdgeListNotEqualValue(t *testing.T) {
	e1 := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(0, 2),
		graphs.NewEdge(1, 2),
	}
	e2 := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(0, 3),
		graphs.NewEdge(1, 2),
	}
	if e1.Equals(e2) {
		t.Errorf("edge lists incorrectly compare equal")
	}
}

func TestEdgeListNotEqualLength(t *testing.T) {
	e1 := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(1, 2),
		graphs.NewEdge(2, 3),
	}
	e2 := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(1, 2),
	}
	if e1.Equals(e2) {
		t.Errorf("edge lists incorrectly compare equal")
	}
}

func TestEdgeListSort(t *testing.T) {
	e1 := graphs.EdgeList{
		graphs.NewEdge(2, 3),
		graphs.NewEdge(0, 1),
		graphs.NewEdge(1, 2),
	}
	e2 := graphs.EdgeList{
		graphs.NewEdge(0, 1),
		graphs.NewEdge(1, 2),
		graphs.NewEdge(2, 3),
	}
	if e1.Equals(e2) {
		t.Errorf("edge lists incorrectly compare equal")
	}
	e1.Sort()
	if !e1.Equals(e2) {
		t.Errorf("edge lists don't compare equal")
	}
}
