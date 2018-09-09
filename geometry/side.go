package geometry

// Side represents a side of a rectangle.
type Side uint8

const (
	// Top represents the top side of a rectangle.
	Top Side = iota
	// Right represents the right side of a rectangle.
	Right
	// Bottom represents the bottom side of a rectangle.
	Bottom
	// Left represents the left side of a rectangle.
	Left
)
