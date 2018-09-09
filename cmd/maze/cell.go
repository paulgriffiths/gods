package main

import "github.com/paulgriffiths/gods/geometry"

// cell implements a cell within a maze.
type cell struct {
	geometry.Point
	size  int
	value byte
}

func (c *cell) corner(cnr geometry.Corner) geometry.Point {
	return c.rect().Corner(cnr)
}

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
// side, to the other cell.
func (c *cell) isAdjacent(other *cell, d geometry.Direction) bool {
	switch d {
	case geometry.North:
		return c.X == other.X && c.Y-other.Y == -1
	case geometry.East:
		return c.X-other.X == 1 && c.Y == other.Y
	case geometry.South:
		return c.X == other.X && c.Y-other.Y == 1
	case geometry.West:
		return c.X-other.X == -1 && c.Y == other.Y
	}
	return false
}
