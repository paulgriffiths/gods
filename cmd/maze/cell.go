package main

import "github.com/paulgriffiths/gods/geometry"

// cell implements a cell within a maze.
type cell struct {
	geometry.Point
	size  int
	value byte
}

// corner returns the coordinates of the specified corner of the cell.
func (c *cell) corner(cnr geometry.Corner) geometry.Point {
	return c.rect().Corner(cnr)
}

// rect returns a rectangle representing the internal cell boundaries.
func (c *cell) rect() geometry.Rect {
	v := geometry.Vector{c.size - 1, c.size - 1}
	topLeft := c.Point.Scale(c.size)
	bottomRight := topLeft.Translate(v)
	return geometry.Rect{topLeft, bottomRight}
}

// hasAttr returns true if the cell has the specified attribute set.
func (c *cell) hasAttr(b byte) bool {
	return c.value&b == b
}

// isAdjacent returns true if the cell is adjacent, on the specified
// side, to the other cell. Returns false if the other cell is nil.
func (c *cell) isAdjacent(other *cell, d geometry.Direction) bool {
	if other == nil {
		return false
	}

	switch d {
	case geometry.North:
		return c.X == other.X && c.Y-other.Y == -1
	case geometry.East:
		return c.X-other.X == 1 && c.Y == other.Y
	case geometry.South:
		return c.X == other.X && c.Y-other.Y == 1
	case geometry.West:
		return c.X-other.X == -1 && c.Y == other.Y
	default:
		panic("undefined direction")
	}
}
