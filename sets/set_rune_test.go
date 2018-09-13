package sets_test

import (
	"github.com/paulgriffiths/gods/sets"
	"testing"
)

var (
	srE    = sets.NewSetRune()
	srA    = sets.NewSetRune('a')
	srB    = sets.NewSetRune('b')
	srC    = sets.NewSetRune('c')
	srAB   = sets.NewSetRune('a', 'b')
	srBA   = sets.NewSetRune('b', 'a')
	srAC   = sets.NewSetRune('a', 'c')
	srBC   = sets.NewSetRune('b', 'c')
	srCD   = sets.NewSetRune('c', 'd')
	srDE   = sets.NewSetRune('d', 'e')
	srABC  = sets.NewSetRune('a', 'b', 'c')
	srBCD  = sets.NewSetRune('b', 'c', 'd')
	srBCDE = sets.NewSetRune('b', 'c', 'd', 'e')
)

func TestSetRuneEquals(t *testing.T) {
	testCases := []struct {
		a, b  sets.SetRune
		equal bool
	}{
		{srE, srE, true}, {srE, srA, false}, {srA, srE, false},
		{srA, srA, true}, {srA, srB, false}, {srB, srA, false},
		{srB, srB, true}, {srAB, srBA, true}, {srAB, srAC, false},
		{srABC, srAB, false},
	}

	for n, tc := range testCases {
		if r := tc.a.Equals(tc.b); r != tc.equal {
			t.Errorf("case %d, got %t, want %t", n+1, r, tc.equal)
		}
	}
}

func TestSetRuneLength(t *testing.T) {
	testCases := []struct {
		values []rune
		length int
	}{
		{[]rune{'a'}, 1}, {[]rune{'a', 'b'}, 2},
		{[]rune{'a', 'a'}, 1}, {[]rune{'a', 'b', 'c'}, 3},
		{[]rune{'a', 'c', 'c'}, 2}, {[]rune{'b', 'b', 'b'}, 1},
	}

	for n, tc := range testCases {
		s := sets.NewSetRune(tc.values...)
		if r := s.Length(); r != tc.length {
			t.Errorf("case %d, got %d, want %d", n+1, r, tc.length)
		}
	}
}

func TestSetRuneContains(t *testing.T) {
	setValues := []rune{'a', 'b', 'c', 'd', 'e'}
	testCases := []struct {
		values   []rune
		contains []bool
	}{
		{[]rune{'a'}, []bool{true, false, false, false, false}},
		{[]rune{'b'}, []bool{false, true, false, false, false}},
		{[]rune{'c'}, []bool{false, false, true, false, false}},
		{[]rune{'d'}, []bool{false, false, false, true, false}},
		{[]rune{'e'}, []bool{false, false, false, false, true}},
		{[]rune{'a', 'b'}, []bool{true, true, false, false, false}},
		{[]rune{'a', 'a', 'b'}, []bool{true, true, false, false, false}},
		{[]rune{'a', 'a', 'a', 'b'}, []bool{true, true, false, false, false}},
		{[]rune{'b', 'c', 'b'}, []bool{false, true, true, false, false}},
		{[]rune{'d', 'd', 'c'}, []bool{false, false, true, true, false}},
		{[]rune{'e', 'd', 'c'}, []bool{false, false, true, true, true}},
		{[]rune{'e', 'c', 'c'}, []bool{false, false, true, false, true}},
	}

	for n, tc := range testCases {
		s := sets.NewSetRune(tc.values...)
		for m, c := range tc.contains {
			if r := s.Contains(setValues[m]); r != c {
				t.Errorf("case (%d,%d), got %t, want %t", n+1, m+1, r, c)
			}
		}
	}
}

func TestSetRuneUnion(t *testing.T) {
	testCases := []struct {
		a, b, u sets.SetRune
	}{
		{srBC, srDE, srBCDE}, {srBC, srCD, srBCD}, {srBC, srBC, srBC},
		{srBC, srB, srBC}, {srBC, srE, srBC}, {srE, srE, srE},
	}

	for n, tc := range testCases {
		if s := tc.a.Union(tc.b); !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.u)
		}
	}
}

func TestSetRuneIntersection(t *testing.T) {
	testCases := []struct {
		a, b, i sets.SetRune
	}{
		{srBC, srDE, srE}, {srBC, srCD, srC}, {srBC, srBC, srBC},
		{srBC, srB, srB}, {srBC, srE, srE}, {srE, srE, srE},
	}

	for n, tc := range testCases {
		if s := tc.a.Intersection(tc.b); !s.Equals(tc.i) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.i)
		}
	}
}
