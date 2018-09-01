package amGraph

import "testing"

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

func TestAMGraph(t *testing.T) {
	g := NewAmGraph(4)
	g.InsertEdge(0, 1)
	g.InsertEdge(0, 2)
	g.InsertEdge(1, 3)
	g.InsertEdge(2, 3)

	a := g.Adjacent(0)
	if !IntSliceEqual(a, []int{1, 2}) {
		t.Errorf("want %v, got %v", []int{1, 2}, a)
	}

	a = g.Adjacent(1)
	if !IntSliceEqual(a, []int{0, 3}) {
		t.Errorf("want %v, got %v", []int{0, 3}, a)
	}

	a = g.Adjacent(2)
	if !IntSliceEqual(a, []int{0, 3}) {
		t.Errorf("want %v, got %v", []int{0, 3}, a)
	}

	a = g.Adjacent(3)
	if !IntSliceEqual(a, []int{1, 2}) {
		t.Errorf("want %v, got %v", []int{1, 2}, a)
	}

}
