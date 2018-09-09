package main

// cell implements a cell within a maze.
type cell struct {
	point
	size  int
	value byte
}

func (c *cell) corner(n cornerValue) point {
	x := c.x * c.size
	y := c.y * c.size
	z := c.size - 1
	if n == topRight || n == bottomRight {
		x += z
	}
	if n == bottomLeft || n == bottomRight {
		y += z
	}
	return point{x, y}
}

func (c *cell) box() box {
	return box{
		point{c.corner(topLeft).x, c.corner(topLeft).y},
		point{c.corner(bottomRight).x, c.corner(bottomRight).y},
	}
}

// hasAttr returns true if the cell has the specified attribute set.
func (c *cell) hasAttr(b byte) bool {
	return c.value&b == b
}

// isAdjacent returns true if the cell is adjacent, on the specified
// side, to the other cell.
func (c *cell) isAdjacent(other *cell, d direction) bool {
	switch d {
	case north:
		return c.x == other.x && c.y-other.y == -1
	case east:
		return c.x-other.x == 1 && c.y == other.y
	case south:
		return c.x == other.x && c.y-other.y == 1
	case west:
		return c.x-other.x == -1 && c.y == other.y
	}
	return false
}
