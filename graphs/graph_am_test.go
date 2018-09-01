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
