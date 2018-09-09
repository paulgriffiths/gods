package geometry

import "testing"

func TestPointEquals(t *testing.T) {
	testCases := []struct {
		a, b   Point
		result bool
	}{
		{Point{0, 0}, Point{0, 0}, true},
		{Point{0, 0}, Point{1, 0}, false},
		{Point{0, 0}, Point{0, 1}, false},
		{Point{0, 0}, Point{1, 1}, false},
	}

	for n, tc := range testCases {
		if result := tc.a.Equals(tc.b); result != tc.result {
			t.Errorf("case %d, got %t, want %t", n+1, result, tc.result)
		}
	}
}

func TestPointScale(t *testing.T) {
	testCases := []struct {
		before, after Point
		factor        int
	}{
		{Point{0, 0}, Point{0, 0}, 0},
		{Point{0, 0}, Point{0, 0}, 1},
		{Point{0, 0}, Point{0, 0}, 2},
		{Point{0, 1}, Point{0, 0}, 0},
		{Point{0, 1}, Point{0, 1}, 1},
		{Point{0, 1}, Point{0, 2}, 2},
		{Point{1, 0}, Point{0, 0}, 0},
		{Point{1, 0}, Point{1, 0}, 1},
		{Point{1, 0}, Point{2, 0}, 2},
		{Point{3, 4}, Point{0, 0}, 0},
		{Point{3, 4}, Point{3, 4}, 1},
		{Point{3, 4}, Point{6, 8}, 2},
		{Point{3, 4}, Point{9, 12}, 3},
		{Point{-3, 4}, Point{0, 0}, 0},
		{Point{-3, 4}, Point{-3, 4}, 1},
		{Point{-3, 4}, Point{-6, 8}, 2},
		{Point{-3, 4}, Point{-9, 12}, 3},
		{Point{3, -4}, Point{-3, 4}, -1},
		{Point{3, -4}, Point{-6, 8}, -2},
		{Point{3, -4}, Point{-9, 12}, -3},
	}

	for n, tc := range testCases {
		if result := tc.before.Scale(tc.factor); !result.Equals(tc.after) {
			t.Errorf("case %d, got %v, want %v", n+1, result, tc.after)
		}
	}
}

func TestPointTranslate(t *testing.T) {
	testCases := []struct {
		before, after Point
		vec           Vector
	}{
		{Point{2, 3}, Point{2, 3}, Vector{0, 0}},
		{Point{2, 3}, Point{8, 3}, Vector{6, 0}},
		{Point{2, 3}, Point{2, 8}, Vector{0, 5}},
		{Point{2, 3}, Point{8, 8}, Vector{6, 5}},
		{Point{2, 3}, Point{-4, 3}, Vector{-6, 0}},
		{Point{2, 3}, Point{2, -2}, Vector{0, -5}},
		{Point{2, 3}, Point{-4, -2}, Vector{-6, -5}},
	}

	for n, tc := range testCases {
		if result := tc.before.Translate(tc.vec); !result.Equals(tc.after) {
			t.Errorf("case %d, got %v, want %v", n+1, result, tc.after)
		}
	}
}
