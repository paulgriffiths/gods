package sets_test

import (
	"github.com/paulgriffiths/gods/sets"
	"testing"
)

type mockType struct {
	n int
	s string
}

func (t mockType) Equals(other mockType) bool {
	return t.n == other.n && t.s == other.s
}

func mockTypeEquals(a, b interface{}) bool {
	return a.(mockType).Equals(b.(mockType))
}

var (
	t0 = mockType{0, "zero"}
	t1 = mockType{1, "one"}
	t2 = mockType{2, "two"}
	t3 = mockType{3, "three"}
	t4 = mockType{4, "four"}
)

func TestSetInterfaceEquals(t *testing.T) {
	testCases := []struct {
		a, b  sets.SetInterface
		equal bool
	}{
		{
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals),
			true,
		},
		{
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals, t0),
			false,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t0),
			sets.NewSetInterface(mockTypeEquals),
			false,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t0),
			sets.NewSetInterface(mockTypeEquals, t0),
			true,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t0),
			sets.NewSetInterface(mockTypeEquals, t1),
			false,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1),
			sets.NewSetInterface(mockTypeEquals, t0),
			false,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1),
			sets.NewSetInterface(mockTypeEquals, t1),
			true,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t0, t1),
			sets.NewSetInterface(mockTypeEquals, t1, t0),
			true,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t0, t1),
			sets.NewSetInterface(mockTypeEquals, t0, t2),
			false,
		},
		{
			sets.NewSetInterface(mockTypeEquals, t0, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t0, t1),
			false,
		},
	}

	for n, tc := range testCases {
		if tc.a.Equals(tc.b) != tc.equal {
			t.Errorf("case %d, got %t, want %t", n+1,
				tc.a.Equals(tc.b), tc.equal)
		}
	}
}

func TestSetInterfaceLength(t *testing.T) {
	testCases := []struct {
		values []mockType
		length int
	}{
		{[]mockType{t1}, 1},
		{[]mockType{t1, t2}, 2},
		{[]mockType{t1, t1}, 1},
		{[]mockType{t1, t2, t3}, 3},
		{[]mockType{t1, t3, t3}, 2},
		{[]mockType{t2, t2, t2}, 1},
	}

	for n, tc := range testCases {
		s := sets.NewSetInterface(mockTypeEquals)
		for _, value := range tc.values {
			s.Insert(value)
		}
		if l := s.Length(); l != tc.length {
			t.Errorf("case %d, got %d, want %d", n+1, l, tc.length)
		}
	}
}

func TestSetInterfaceContains(t *testing.T) {
	setValues := []mockType{t0, t1, t2, t3, t4}
	testCases := []struct {
		values   []mockType
		contains []bool
	}{
		{[]mockType{t0}, []bool{true, false, false, false, false}},
		{[]mockType{t1}, []bool{false, true, false, false, false}},
		{[]mockType{t2}, []bool{false, false, true, false, false}},
		{[]mockType{t3}, []bool{false, false, false, true, false}},
		{[]mockType{t4}, []bool{false, false, false, false, true}},
		{[]mockType{t0, t1}, []bool{true, true, false, false, false}},
		{[]mockType{t0, t0, t1}, []bool{true, true, false, false, false}},
		{[]mockType{t0, t0, t0, t1}, []bool{true, true, false, false, false}},
		{[]mockType{t1, t2, t1}, []bool{false, true, true, false, false}},
		{[]mockType{t3, t3, t2}, []bool{false, false, true, true, false}},
		{[]mockType{t4, t3, t2}, []bool{false, false, true, true, true}},
		{[]mockType{t4, t2, t2}, []bool{false, false, true, false, true}},
	}

	for n, tc := range testCases {
		s := sets.NewSetInterface(mockTypeEquals)
		for _, value := range tc.values {
			s.Insert(value)
		}
		for m, c := range tc.contains {
			if r := s.Contains(setValues[m]); r != c {
				t.Errorf("case (%d,%d), got %t, want %t", n, m, r, c)
			}
		}
	}
}

func TestSetInterfaceUnion(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInterface
	}{
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t3, t4),
			sets.NewSetInterface(mockTypeEquals, t1, t2, t3, t4),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t2, t3),
			sets.NewSetInterface(mockTypeEquals, t1, t2, t3),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t1, t2),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t1),
			sets.NewSetInterface(mockTypeEquals, t1, t2),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals, t1, t2),
		},
		{
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals),
		},
	}

	for n, tc := range testCases {
		s := tc.a.Union(tc.b)
		if !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n, s, tc.u)
		}
	}
}

func TestSetInterfaceIntersection(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInterface
	}{
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t3, t4),
			sets.NewSetInterface(mockTypeEquals),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t2, t3),
			sets.NewSetInterface(mockTypeEquals, t2),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t1, t2),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals, t1),
			sets.NewSetInterface(mockTypeEquals, t1),
		},
		{
			sets.NewSetInterface(mockTypeEquals, t1, t2),
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals),
		},
		{
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals),
			sets.NewSetInterface(mockTypeEquals),
		},
	}

	for n, tc := range testCases {
		s := tc.a.Intersection(tc.b)
		if !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n, s, tc.u)
		}
	}
}
