package sets_test

import (
	"github.com/paulgriffiths/gods/sets"
	"testing"
)

var (
	se    = sets.NewSetInt()
	s0    = sets.NewSetInt(0)
	s1    = sets.NewSetInt(1)
	s2    = sets.NewSetInt(2)
	s01   = sets.NewSetInt(0, 1)
	s10   = sets.NewSetInt(1, 0)
	s02   = sets.NewSetInt(0, 2)
	s12   = sets.NewSetInt(1, 2)
	s23   = sets.NewSetInt(2, 3)
	s34   = sets.NewSetInt(3, 4)
	s012  = sets.NewSetInt(0, 1, 2)
	s123  = sets.NewSetInt(1, 2, 3)
	s1234 = sets.NewSetInt(1, 2, 3, 4)
)

func TestSetIntEquals(t *testing.T) {
	testCases := []struct {
		a, b  sets.SetInt
		equal bool
	}{
		{se, se, true},
		{se, s0, false},
		{s0, se, false},
		{s0, s0, true},
		{s0, s1, false},
		{s1, s0, false},
		{s1, s1, true},
		{s01, s10, true},
		{s01, s02, false},
		{s012, s01, false},
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
		s := sets.NewSetInt(tc.values...)
		if r := s.Length(); r != tc.length {
			t.Errorf("case %d, got %d, want %d", n+1, r, tc.length)
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
		s := sets.NewSetInt(tc.values...)
		for m, c := range tc.contains {
			if r := s.Contains(m); r != c {
				t.Errorf("case (%d,%d), got %t, want %t", n+1, m+1, r, c)
			}
		}
	}
}

func TestSetUnion(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInt
	}{
		{s12, s34, s1234},
		{s12, s23, s123},
		{s12, s12, s12},
		{s12, s1, s12},
		{s12, se, s12},
		{se, se, se},
	}

	for n, tc := range testCases {
		if s := tc.a.Union(tc.b); !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.u)
		}
	}
}

func TestSetIntersection(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInt
	}{
		{s12, s34, se},
		{s12, s23, s2},
		{s12, s12, s12},
		{s12, s1, s1},
		{s12, se, se},
		{se, se, se},
	}

	for n, tc := range testCases {
		if s := tc.a.Intersection(tc.b); !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.u)
		}
	}
}
