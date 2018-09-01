package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestVertexListLength(t *testing.T) {
	v := graphs.VertexList{0, 1, 2}
	if v.Len() != 3 {
		t.Errorf("got %d, want %d", v.Len(), 3)
	}
}

func TestVertexListEqual(t *testing.T) {
	v1 := graphs.VertexList{0, 1, 2}
	v2 := graphs.VertexList{0, 1, 2}
	if !v1.Equals(v2) {
		t.Errorf("vertex lists don't compare equal")
	}
}

func TestVertexListNotEqualValue(t *testing.T) {
	v1 := graphs.VertexList{0, 1, 2}
	v2 := graphs.VertexList{0, 1, 3}
	if v1.Equals(v2) {
		t.Errorf("vertex lists incorrectly compare equal")
	}
}

func TestVertexListNotEqualLength(t *testing.T) {
	v1 := graphs.VertexList{0, 1, 2}
	v2 := graphs.VertexList{0, 1}
	if v1.Equals(v2) {
		t.Errorf("vertex lists incorrectly compare equal")
	}
}

func TestVertexListSort(t *testing.T) {
	v1 := graphs.VertexList{2, 0, 3, 1}
	v2 := graphs.VertexList{0, 1, 2, 3}
	if v1.Equals(v2) {
		t.Errorf("vertex lists incorrectly compare equal")
	}
	v1.Sort()
	if !v1.Equals(v2) {
		t.Errorf("vertex lists don't compare equal")
	}
}
