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
	m0 = mockType{0, "zero"}
	m1 = mockType{1, "one"}
	m2 = mockType{2, "two"}
	m3 = mockType{3, "three"}
	m4 = mockType{4, "four"}
)

var (
	smE    = sets.NewSetInterface(mockTypeEquals)
	sm0    = sets.NewSetInterface(mockTypeEquals, m0)
	sm1    = sets.NewSetInterface(mockTypeEquals, m1)
	sm2    = sets.NewSetInterface(mockTypeEquals, m2)
	sm01   = sets.NewSetInterface(mockTypeEquals, m0, m1)
	sm10   = sets.NewSetInterface(mockTypeEquals, m1, m0)
	sm02   = sets.NewSetInterface(mockTypeEquals, m0, m2)
	sm12   = sets.NewSetInterface(mockTypeEquals, m1, m2)
	sm23   = sets.NewSetInterface(mockTypeEquals, m2, m3)
	sm34   = sets.NewSetInterface(mockTypeEquals, m3, m4)
	sm012  = sets.NewSetInterface(mockTypeEquals, m0, m1, m2)
	sm123  = sets.NewSetInterface(mockTypeEquals, m1, m2, m3)
	sm1234 = sets.NewSetInterface(mockTypeEquals, m1, m2, m3, m4)
)

func TestSetInterfaceEquals(t *testing.T) {
	testCases := []struct {
		a, b  sets.SetInterface
		equal bool
	}{
		{smE, smE, true}, {smE, sm0, false}, {sm0, smE, false},
		{sm0, sm0, true}, {sm0, sm1, false}, {sm1, sm0, false},
		{sm1, sm1, true}, {sm01, sm10, true}, {sm01, sm02, false},
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
		{[]mockType{m1}, 1},
		{[]mockType{m1, m2}, 2},
		{[]mockType{m1, m1}, 1},
		{[]mockType{m1, m2, m3}, 3},
		{[]mockType{m1, m3, m3}, 2},
		{[]mockType{m2, m2, m2}, 1},
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
	setValues := []mockType{m0, m1, m2, m3, m4}
	testCases := []struct {
		values   []mockType
		contains []bool
	}{
		{[]mockType{m0}, []bool{true, false, false, false, false}},
		{[]mockType{m1}, []bool{false, true, false, false, false}},
		{[]mockType{m2}, []bool{false, false, true, false, false}},
		{[]mockType{m3}, []bool{false, false, false, true, false}},
		{[]mockType{m4}, []bool{false, false, false, false, true}},
		{[]mockType{m0, m1}, []bool{true, true, false, false, false}},
		{[]mockType{m0, m0, m1}, []bool{true, true, false, false, false}},
		{[]mockType{m0, m0, m0, m1}, []bool{true, true, false, false, false}},
		{[]mockType{m1, m2, m1}, []bool{false, true, true, false, false}},
		{[]mockType{m3, m3, m2}, []bool{false, false, true, true, false}},
		{[]mockType{m4, m3, m2}, []bool{false, false, true, true, true}},
		{[]mockType{m4, m2, m2}, []bool{false, false, true, false, true}},
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
		{sm12, sm34, sm1234}, {sm12, sm23, sm123}, {sm12, sm12, sm12},
		{sm12, sm1, sm12}, {sm12, smE, sm12}, {smE, smE, smE},
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
		{sm12, sm34, smE}, {sm12, sm23, sm2}, {sm12, sm12, sm12},
		{sm12, sm1, sm1}, {sm12, smE, smE}, {smE, smE, smE},
	}

	for n, tc := range testCases {
		if s := tc.a.Intersection(tc.b); !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.u)
		}
	}
}
