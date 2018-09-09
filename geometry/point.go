package geometry

// Point implements a two-dimensional point value.
type Point struct {
	X, Y int
}

// Equals tests if two points are equal.
func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

// Scale returns a new point with the X, Y coordinates multiplied
// by the provided scale factor.
func (p Point) Scale(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// Translate returns a new point with the coordinates increased
// by the specified amounts.
func (p Point) Translate(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y}
}
