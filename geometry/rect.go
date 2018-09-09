package geometry

// Rect implements a rectangle.
type Rect struct {
	TopLeft, BottomRight Point
}

// Equals tests if two rectangles are equal.
func (r Rect) Equals(other Rect) bool {
	return r.TopLeft == other.TopLeft && r.BottomRight == other.BottomRight
}

// Scale returns a new rectangle with the X, Y coordinates multiplied
// by the provided scale factor.
func (r Rect) Scale(factor int) Rect {
	return Rect{r.TopLeft.Scale(factor), r.BottomRight.Scale(factor)}
}

// Translate returns a new rectangle with the coordinates increased
// by the specified amounts.
func (r Rect) Translate(v Vector) Rect {
	return Rect{r.TopLeft.Translate(v), r.BottomRight.Translate(v)}
}

// Inset returns a rectangle inset from the original dimensions
// by the specified inset values.
// NOTE: the directions imply a y-axis increasing in the down direction.
func (r Rect) Inset(top, right, bottom, left int) Rect {
	return Rect{
		Point{r.TopLeft.X + left, r.TopLeft.Y + top},
		Point{r.BottomRight.X - right, r.BottomRight.Y - bottom},
	}
}

// Corner returns the coordinates of the specified corner
// of the rectangle.
func (r Rect) Corner(c Corner) Point {
	switch c {
	case TopLeft:
		return r.TopLeft
	case TopRight:
		return Point{r.BottomRight.X, r.TopLeft.Y}
	case BottomLeft:
		return Point{r.TopLeft.X, r.BottomRight.Y}
	case BottomRight:
		return r.BottomRight
	default:
		panic("undefined corner value")
	}
}
