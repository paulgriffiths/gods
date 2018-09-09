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

var (
	sme    = sets.NewSetInterface(mockTypeEquals)
	sm0    = sets.NewSetInterface(mockTypeEquals, t0)
	sm1    = sets.NewSetInterface(mockTypeEquals, t1)
	sm2    = sets.NewSetInterface(mockTypeEquals, t2)
	sm01   = sets.NewSetInterface(mockTypeEquals, t0, t1)
	sm10   = sets.NewSetInterface(mockTypeEquals, t1, t0)
	sm02   = sets.NewSetInterface(mockTypeEquals, t0, t2)
	sm12   = sets.NewSetInterface(mockTypeEquals, t1, t2)
	sm23   = sets.NewSetInterface(mockTypeEquals, t2, t3)
	sm34   = sets.NewSetInterface(mockTypeEquals, t3, t4)
	sm012  = sets.NewSetInterface(mockTypeEquals, t0, t1, t2)
	sm123  = sets.NewSetInterface(mockTypeEquals, t1, t2, t3)
	sm1234 = sets.NewSetInterface(mockTypeEquals, t1, t2, t3, t4)
)

func TestSetInterfaceEquals(t *testing.T) {
	testCases := []struct {
		a, b  sets.SetInterface
		equal bool
	}{
		{sme, sme, true},
		{sme, sm0, false},
		{sm0, sme, false},
		{sm0, sm0, true},
		{sm0, sm1, false},
		{sm1, sm0, false},
		{sm1, sm1, true},
		{sm01, sm10, true},
		{sm01, sm02, false},
		{sm012, sm01, false},
	}

	for n, tc := range testCases {
		if r := tc.a.Equals(tc.b); r != tc.equal {
			t.Errorf("case %d, got %t, want %t", n+1, r, tc.equal)
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
				t.Errorf("case (%d,%d), got %t, want %t", n+1, m, r, c)
			}
		}
	}
}

func TestSetInterfaceUnion(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInterface
	}{
		{sm12, sm34, sm1234},
		{sm12, sm23, sm123},
		{sm12, sm12, sm12},
		{sm12, sm1, sm12},
		{sm12, sme, sm12},
		{sme, sme, sme},
	}

	for n, tc := range testCases {
		if s := tc.a.Union(tc.b); !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.u)
		}
	}
}

func TestSetInterfaceIntersection(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetInterface
	}{
		{sm12, sm34, sme},
		{sm12, sm23, sm2},
		{sm12, sm12, sm12},
		{sm12, sm1, sm1},
		{sm12, sme, sme},
		{sme, sme, sme},
	}

	for n, tc := range testCases {
		if s := tc.a.Intersection(tc.b); !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.u)
		}
	}
}
