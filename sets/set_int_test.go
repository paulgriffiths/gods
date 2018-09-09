package sets_test

import (
	"github.com/paulgriffiths/gods/sets"
	"testing"
)

func TestSetIntEquals(t *testing.T) {
	testCases := []struct {
		a, b  sets.SetInt
		equal bool
	}{
		{sets.NewSetInt(), sets.NewSetInt(), true},
		{sets.NewSetInt(), sets.NewSetInt(0), false},
		{sets.NewSetInt(0), sets.NewSetInt(), false},
		{sets.NewSetInt(0), sets.NewSetInt(0), true},
		{sets.NewSetInt(0), sets.NewSetInt(1), false},
		{sets.NewSetInt(1), sets.NewSetInt(0), false},
		{sets.NewSetInt(1), sets.NewSetInt(1), true},
		{sets.NewSetInt(0, 1), sets.NewSetInt(1, 0), true},
		{sets.NewSetInt(0, 1), sets.NewSetInt(0, 2), false},
		{sets.NewSetInt(0, 1, 2), sets.NewSetInt(0, 1), false},
	}

	for n, tc := range testCases {
		if r := tc.a.Equals(tc.b); r != tc.equal {
			t.Errorf("case %d, got %t, want %t", n+1, r, tc.equal)
		}
	}
}

func TestSetIntLength(t *testing.T) {
	testCases := []struct {
		values []int
		length int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 3, 3}, 2},
		{[]int{2, 2, 2}, 1},
	}

	for n, tc := range testCases {
		s := sets.NewSetInt()
		for _, value := range tc.values {
			s.Insert(value)
		}
		if r := s.Length(); r != tc.length {
			t.Errorf("case %d, got %d, want %d", n, r, tc.length)
		}
	}
}

func TestSetContains(t *testing.T) {
	testCases := []struct {
		values   []int
		contains []bool
	}{
		{[]int{0}, []bool{true, false, false, false, false}},
		{[]int{1}, []bool{false, true, false, false, false}},
		{[]int{2}, []bool{false, false, true, false, false}},
		{[]int{3}, []bool{false, false, false, true, false}},
		{[]int{4}, []bool{false, false, false, false, true}},
		{[]int{0, 1}, []bool{true, true, false, false, false}},
		{[]int{0, 0, 1}, []bool{true, true, false, false, false}},
		{[]int{0, 0, 0, 1}, []bool{true, true, false, false, false}},
		{[]int{1, 2, 1}, []bool{false, true, true, false, false}},
		{[]int{3, 3, 2}, []bool{false, false, true, true, false}},
		{[]int{4, 3, 2}, []bool{false, false, true, true, true}},
		{[]int{4, 2, 2}, []bool{false, false, true, false, true}},
	}

	for n, tc := range testCases {
		s := sets.NewSetInt()
		for _, value := range tc.values {
			s.Insert(value)
		}
		for m, c := range tc.contains {
			if r := s.Contains(m); r != c {
				t.Errorf("case (%d,%d), got %t, want %t", n, m, r, c)
			}
		}
	}
}

func TestSetUnion(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInt
	}{
		{sets.NewSetInt(1, 2), sets.NewSetInt(3, 4), sets.NewSetInt(1, 2, 3, 4)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(2, 3), sets.NewSetInt(1, 2, 3)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(1, 2), sets.NewSetInt(1, 2)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(1), sets.NewSetInt(1, 2)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(), sets.NewSetInt(1, 2)},
		{sets.NewSetInt(), sets.NewSetInt(), sets.NewSetInt()},
	}

	for n, tc := range testCases {
		s := tc.a.Union(tc.b)
		if !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n, s, tc.u)
		}
	}
}

func TestSetIntersection(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInt
	}{
		{sets.NewSetInt(1, 2), sets.NewSetInt(3, 4), sets.NewSetInt()},
		{sets.NewSetInt(1, 2), sets.NewSetInt(2, 3), sets.NewSetInt(2)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(1, 2), sets.NewSetInt(1, 2)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(1), sets.NewSetInt(1)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(), sets.NewSetInt()},
		{sets.NewSetInt(), sets.NewSetInt(), sets.NewSetInt()},
	}

	for n, tc := range testCases {
		s := tc.a.Intersection(tc.b)
		if !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n, s, tc.u)
		}
	}
}
