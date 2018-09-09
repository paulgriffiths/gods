package geometry

// Corner represents the corner of a rectangle.
type Corner uint8

const (
	// TopLeft represents the top-left corner of a rectangle.
	TopLeft Corner = iota
	// TopRight represents the top-right side of a rectangle.
	TopRight
	// BottomLeft represents the bottom-left side of a rectangle.
	BottomLeft
	// BottomRight represents the bottom-right side of a rectangle.
	BottomRight
)
