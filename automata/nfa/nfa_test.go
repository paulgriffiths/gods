/*
NFA models used for testing taken from "Introduction to the Theory of
Computation", Third Edition, Michael Sipser. (C) 2013 Cengage Learning,
and from "Compilers", Second Edition, Aho, Lam, Sethi & Ullman, (C) 2007
Pearson Education, Inc.
*/
package nfa_test

import (
	_ "fmt"
	"github.com/paulgriffiths/gods/automata/nfa"
	"github.com/paulgriffiths/gods/sets"
	"testing"
)

// NFA from Introduction to the Theory of Computation, figure 1.34.
func TestNfaEclosure1(t *testing.T) {
	n := nfa.Nfa{
		6,
		[]rune{'0', '1'},
		[]map[rune]sets.SetInt{
			{0: {1, 2}},
			{'0': {2}},
			{'0': {1}},
			{'0': {4}},
			{'0': {5}},
			{'0': {3}},
		},
		0,
		sets.SetInt{1, 3},
	}

	testCases := []sets.SetInt{
		sets.SetInt{0, 1, 2},
		sets.SetInt{1},
		sets.SetInt{2},
		sets.SetInt{3},
		sets.SetInt{4},
		sets.SetInt{5},
	}

	for i, tc := range testCases {
		s := n.Eclosure(sets.SetInt{i})
		if !s.Equals(tc) {
			t.Errorf("case %d, got %v, want %v", i+1, s, tc)
		}
	}
}

// NFA from Compilers, figure 3.34.
func TestNfaEclosure2(t *testing.T) {
	n := nfa.Nfa{
		11,
		[]rune{'a', 'b'},
		[]map[rune]sets.SetInt{
			{0: {1, 7}},
			{0: {2, 4}},
			{'a': {3}},
			{0: {6}},
			{'b': {5}},
			{0: {6}},
			{0: {1, 7}},
			{'a': {8}},
			{'b': {9}},
			{'b': {10}},
			{},
		},
		0,
		sets.SetInt{10},
	}

	testCases := []sets.SetInt{
		sets.SetInt{0, 1, 2, 4, 7},
		sets.SetInt{1, 2, 4},
		sets.SetInt{2},
		sets.SetInt{1, 2, 3, 4, 6, 7},
		sets.SetInt{4},
		sets.SetInt{1, 2, 4, 5, 6, 7},
		sets.SetInt{1, 2, 4, 6, 7},
		sets.SetInt{7},
		sets.SetInt{8},
		sets.SetInt{9},
		sets.SetInt{10},
	}

	for i, tc := range testCases {
		s := n.Eclosure(sets.SetInt{i})
		if !s.Equals(tc) {
			t.Errorf("case %d, got %v, want %v", i+1, s, tc)
		}
	}
}

// Nfa1 recognizes (a|b)*abb.
// Compilers, figure 3.24.
func TestNfa1(t *testing.T) {
	n := nfa.Nfa{
		4,
		[]rune{'a', 'b'},
		[]map[rune]sets.SetInt{
			{'a': {0, 1}, 'b': {0}},
			{'b': {2}},
			{'b': {3}},
			{},
		},
		0,
		sets.SetInt{3},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"a", false},
		{"b", false},
		{"aa", false},
		{"ab", false},
		{"ba", false},
		{"bb", false},
		{"aaa", false},
		{"aab", false},
		{"aba", false},
		{"abb", true},
		{"baa", false},
		{"bab", false},
		{"bba", false},
		{"bbb", false},
		{"aabb", true},
		{"babb", true},
		{"aaabb", true},
		{"ababb", true},
		{"baabb", true},
		{"bbabb", true},
		{"aaabba", false},
		{"ababba", false},
		{"baabba", false},
		{"bbabba", false},
		{"aaabbb", false},
		{"ababbb", false},
		{"baabbb", false},
		{"bbabbb", false},
		{"aaaaaaaabb", true},
		{"bbbbbbbabb", true},
		{"abababababb", true},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Nfa2 recognizes aa*|bb*
// Compilers, figure 3.26.
func TestNfa2(t *testing.T) {
	n := nfa.Nfa{
		5,
		[]rune{'a', 'b'},
		[]map[rune]sets.SetInt{
			{0: {1, 3}},
			{'a': {2}},
			{'a': {2}},
			{'b': {4}},
			{'b': {4}},
		},
		0,
		sets.SetInt{2, 4},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"a", true},
		{"b", true},
		{"aa", true},
		{"ab", false},
		{"ba", false},
		{"bb", true},
		{"aaa", true},
		{"aab", false},
		{"aba", false},
		{"abb", false},
		{"baa", false},
		{"bab", false},
		{"bba", false},
		{"bbb", true},
		{"abab", false},
		{"abba", false},
		{"baba", false},
		{"baab", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Nfa3 recognizes strings containing at least 2 'a's and
// ending with a 'b'.
// Compilers, figure 3.29.
func TestNfa3(t *testing.T) {
	n := nfa.Nfa{
		4,
		[]rune{'a', 'b'},
		[]map[rune]sets.SetInt{
			{'a': {0, 1}, 'b': {0}},
			{'a': {1, 2}, 'b': {1}},
			{0: {0}, 'a': {2}, 'b': {2, 3}},
			{},
		},
		0,
		sets.SetInt{3},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"a", false},
		{"b", false},
		{"aa", false},
		{"ab", false},
		{"ba", false},
		{"bb", false},
		{"aaa", false},
		{"aab", true},
		{"aba", false},
		{"abb", false},
		{"baa", false},
		{"bab", false},
		{"bba", false},
		{"bbb", false},
		{"abab", true},
		{"abba", false},
		{"baba", false},
		{"baab", true},
		{"babbbbbbbbbbbbb", false},
		{"babbbbbbbbbbbba", false},
		{"babbbbbbbbbbbbab", true},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Nfa4 recognizes (a|b)*.
// Compilers, figure 3.30.
func TestNfa4(t *testing.T) {
	n := nfa.Nfa{
		4,
		[]rune{'a', 'b'},
		[]map[rune]sets.SetInt{
			{0: {3}, 'a': {1}},
			{0: {0}, 'b': {2}},
			{0: {1}, 'b': {3}},
			{0: {2}, 'a': {0}},
		},
		0,
		sets.SetInt{3},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"a", true},
		{"b", true},
		{"aa", true},
		{"ab", true},
		{"ba", true},
		{"bb", true},
		{"aaa", true},
		{"aab", true},
		{"aba", true},
		{"abb", true},
		{"baa", true},
		{"bab", true},
		{"bba", true},
		{"bbb", true},
		{"abab", true},
		{"abba", true},
		{"baba", true},
		{"baab", true},
		{"abbbbbbbbbbbbb", true},
		{"abbbbbbbbbbbba", true},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Nfa5 recognizes strings that have either 101 or 11 as a substring.
// Introduction to the Theory of Computation, figure 1.27.
func TestNfa5(t *testing.T) {
	n := nfa.Nfa{
		4,
		[]rune{'0', '1'},
		[]map[rune]sets.SetInt{
			{'0': {0}, '1': {0, 1}},
			{0: {2}, '0': {2}},
			{'1': {3}},
			{'0': {3}, '1': {3}},
		},
		0,
		sets.SetInt{3},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"0", false},
		{"1", false},
		{"00", false},
		{"01", false},
		{"10", false},
		{"11", true},
		{"000", false},
		{"001", false},
		{"010", false},
		{"011", true},
		{"100", false},
		{"101", true},
		{"110", true},
		{"111", true},
		{"001001001001001001001", false},
		{"0010010010011001001001", true},
		{"0010010010100001001001", true},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Nfa6 recognizes strings that have a '1' in the third position
// from the end.
// Introduction to the Theory of Computation, figure 1.31.
func TestNfa6(t *testing.T) {
	n := nfa.Nfa{
		4,
		[]rune{'0', '1'},
		[]map[rune]sets.SetInt{
			{'0': {0}, '1': {0, 1}},
			{'0': {2}, '1': {2}},
			{'0': {3}, '1': {3}},
			{},
		},
		0,
		sets.SetInt{3},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"0", false},
		{"1", false},
		{"00", false},
		{"01", false},
		{"10", false},
		{"11", false},
		{"000", false},
		{"001", false},
		{"010", false},
		{"011", false},
		{"100", true},
		{"101", true},
		{"110", true},
		{"111", true},
		{"10101010101010101010101010101", true},
		{"10101010101010101010101010001", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Nfa7 recognizes strings of '0's where the total number of zeroes
// is a multiple of 2 or 3.
// Introduction to the Theory of Computation, figure 1.34.
func TestNfa7(t *testing.T) {
	n := nfa.Nfa{
		6,
		[]rune{'0', '1'},
		[]map[rune]sets.SetInt{
			{0: {1, 3}},
			{'0': {2}},
			{'0': {1}},
			{'0': {4}},
			{'0': {5}},
			{'0': {3}},
		},
		0,
		sets.SetInt{1, 3},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"0", false},
		{"00", true},
		{"000", true},
		{"0000", true},
		{"00000", false},
		{"000000", true},
		{"0000000", false},
		{"00000000", true},
		{"0000000000", true},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Nfa8 recognizes (a|bba|baa|baa*ba)*
// Introduction to the Theory of Computation, figure 1.36.
func TestNfa8(t *testing.T) {
	n := nfa.Nfa{
		3,
		[]rune{'a', 'b'},
		[]map[rune]sets.SetInt{
			{0: {2}, 'b': {1}},
			{'a': {1, 2}, 'b': {2}},
			{'a': {0}},
		},
		0,
		sets.SetInt{0},
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"a", true},
		{"b", false},
		{"aa", true},
		{"ab", false},
		{"ba", false},
		{"bb", false},
		{"aaa", true},
		{"aab", false},
		{"aba", false},
		{"abb", false},
		{"baa", true},
		{"bab", false},
		{"bba", true},
		{"bbb", false},
		{"baba", true},
		{"babba", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}
