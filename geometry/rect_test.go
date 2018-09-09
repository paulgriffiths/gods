package geometry

import "testing"

var (
	r0000 = Rect{Point{0, 0}, Point{0, 0}}
	r0012 = Rect{Point{0, 0}, Point{1, 2}}
	r0024 = Rect{Point{0, 0}, Point{2, 4}}
	r0036 = Rect{Point{0, 0}, Point{3, 6}}
	r1123 = Rect{Point{1, 1}, Point{2, 3}}
	r1200 = Rect{Point{1, 2}, Point{0, 0}}
	r1212 = Rect{Point{1, 2}, Point{1, 2}}
	r1301 = Rect{Point{1, 3}, Point{0, 1}}
	r1402 = Rect{Point{1, 4}, Point{0, 2}}
	r2210 = Rect{Point{2, 2}, Point{1, 0}}
	r2234 = Rect{Point{2, 2}, Point{3, 4}}
	r2400 = Rect{Point{2, 4}, Point{0, 0}}
	r3220 = Rect{Point{3, 2}, Point{2, 0}}
	r3600 = Rect{Point{3, 6}, Point{0, 0}}
	r0055 = Rect{Point{0, 0}, Point{5, 5}}
	r1541 = Rect{Point{1, 5}, Point{4, 1}}
)

func TestRectEquals(t *testing.T) {
	testCases := []struct {
		a, b   Rect
		result bool
	}{
		{r1200, r1200, true}, {r1200, r1212, false},
		{r1212, r1200, false},
	}

	for n, tc := range testCases {
		if result := tc.a.Equals(tc.b); result != tc.result {
			t.Errorf("case %d, got %t, want %t", n+1, result, tc.result)
		}
	}
}

func TestRectScale(t *testing.T) {
	testCases := []struct {
		before, after Rect
		factor        int
	}{
		{r0000, r0000, 0}, {r0000, r0000, 1},
		{r0000, r0000, 2}, {r0000, r0000, 3},
		{r0012, r0000, 0}, {r0012, r0012, 1},
		{r0012, r0024, 2}, {r0012, r0036, 3},
		{r1200, r0000, 0}, {r1200, r1200, 1},
		{r1200, r2400, 2}, {r1200, r3600, 3},
	}

	for n, tc := range testCases {
		if result := tc.before.Scale(tc.factor); !result.Equals(tc.after) {
			t.Errorf("case %d, got %v, want %v", n+1, result, tc.after)
		}
	}
}

func TestRectTranslate(t *testing.T) {
	testCases := []struct {
		before, after Rect
		vec           Vector
	}{
		{r0012, r0012, Vector{0, 0}}, {r0012, r1123, Vector{1, 1}},
		{r0012, r2234, Vector{2, 2}}, {r2234, r1123, Vector{-1, -1}},
		{r2234, r0012, Vector{-2, -2}}, {r1200, r2210, Vector{1, 0}},
		{r1200, r3220, Vector{2, 0}}, {r3220, r2210, Vector{-1, 0}},
		{r3220, r1200, Vector{-2, 0}}, {r1200, r1301, Vector{0, 1}},
		{r1200, r1402, Vector{0, 2}}, {r1402, r1301, Vector{0, -1}},
		{r1402, r1200, Vector{0, -2}},
	}

	for n, tc := range testCases {
		if result := tc.before.Translate(tc.vec); !result.Equals(tc.after) {
			t.Errorf("case %d, got %v, want %v", n+1, result, tc.after)
		}
	}
}

func TestRectInset(t *testing.T) {
	testCases := []struct {
		before, after Rect
		t, r, b, l    int
	}{
		{r0055, r0012, 0, 4, 3, 0},
		{r0055, r2234, 2, 2, 1, 2},
		{r0055, r1123, 1, 3, 2, 1},
	}

	for n, tc := range testCases {
		result := tc.before.Inset(tc.t, tc.r, tc.b, tc.l)
		if !result.Equals(tc.after) {
			t.Errorf("case %d, got %v, want %v", n+1, result, tc.after)
		}
	}
}

func TestRectCorner(t *testing.T) {
	testCases := []struct {
		r Rect
		c Corner
		p Point
	}{
		{r1541, TopLeft, Point{1, 5}},
		{r1541, TopRight, Point{4, 5}},
		{r1541, BottomLeft, Point{1, 1}},
		{r1541, BottomRight, Point{4, 1}},
		{r0055, TopLeft, Point{0, 0}},
		{r0055, TopRight, Point{5, 0}},
		{r0055, BottomRight, Point{5, 5}},
		{r0055, BottomLeft, Point{0, 5}},
	}

	for n, tc := range testCases {
		if result := tc.r.Corner(tc.c); !result.Equals(tc.p) {
			t.Errorf("case %d, got %v, want %v", n+1, result, tc.p)
		}
	}
}
