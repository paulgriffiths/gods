package graphs_test

import (
	"github.com/paulgriffiths/gods/graphs"
	"testing"
)

func TestEdgeEquals(t *testing.T) {
	testCases := []struct {
		edges  [2]graphs.Edge
		result bool
	}{
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(1, 2)}, true},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(2, 1)}, true},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(1, 3)}, false},
		{[2]graphs.Edge{graphs.NewEdge(3, 2), graphs.NewEdge(1, 2)}, false},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(3, 4)}, false},
	}
	for n, testCase := range testCases {
		result := testCase.edges[0].Equals(testCase.edges[1])
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v",
				n+1, result, testCase.result)
		}
	}
}

func TestEdgeLess(t *testing.T) {
	testCases := []struct {
		edges  [2]graphs.Edge
		result bool
	}{
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(0, 1)}, false},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(0, 2)}, false},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(0, 3)}, false},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(1, 2)}, false},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(1, 3)}, true},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(1, 4)}, true},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(2, 3)}, true},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(2, 4)}, true},
		{[2]graphs.Edge{graphs.NewEdge(1, 2), graphs.NewEdge(2, 5)}, true},
	}
	for n, testCase := range testCases {
		result := testCase.edges[0].Less(testCase.edges[1])
		if result != testCase.result {
			t.Errorf("case %d, got %v, want %v",
				n+1, result, testCase.result)
		}
	}
}
