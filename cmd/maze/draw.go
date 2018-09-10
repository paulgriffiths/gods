package main

import (
	"github.com/paulgriffiths/gods/geometry"
	"image"
)

// drawLine draws a line of the specified color between two points.
// The points must form either a vertical or a horizontal line.
func drawLine(img *image.Paletted, a, b geometry.Point, color uint8) {
	if a.Equals(b) {
		return
	}

	var v geometry.Vector
	switch {
	case a.X == b.X && a.Y < b.Y:
		v = geometry.Vector{0, 1}
	case a.X == b.X && a.Y > b.Y:
		v = geometry.Vector{0, -1}
	case a.X < b.X && a.Y == b.Y:
		v = geometry.Vector{1, 0}
	case a.X > b.X && a.Y == b.Y:
		v = geometry.Vector{-1, 0}
	default:
		panic("points are not in a straight line")
	}

	for p := a; ; p = p.Translate(v) {
		img.SetColorIndex(p.X, p.Y, color)
		if p.Equals(b) {
			break
		}
	}
}

// drawCellWalls draws the walls of a cell in the specified color.
func drawCellWalls(img *image.Paletted, c *cell, color uint8) {
	tl := c.corner(geometry.TopLeft)
	tr := c.corner(geometry.TopRight)
	bl := c.corner(geometry.BottomLeft)
	br := c.corner(geometry.BottomRight)

	if c.hasAttr(northWall) {
		drawLine(img, tl, tr, color)
	}
	if c.hasAttr(eastWall) {
		drawLine(img, tr, br, color)
	}
	if c.hasAttr(southWall) {
		drawLine(img, bl, br, color)
	}
	if c.hasAttr(westWall) {
		drawLine(img, tl, bl, color)
	}
}

// drawBox draws a solid box in the specified color.
func drawBox(img *image.Paletted, r geometry.Rect, color uint8) {
	for x := r.TopLeft.X; x <= r.BottomRight.X; x++ {
		for y := r.TopLeft.Y; y <= r.BottomRight.Y; y++ {
			img.SetColorIndex(x, y, color)
		}
	}
}

// drawPath draws a path marker in the center of a cell.
func drawPath(img *image.Paletted, c, from, to *cell, color uint8) {
	off := c.size / 3

	drawBox(img, c.rect().Inset(off, off, off, off), color)

	if c.isAdjacent(from, geometry.North) ||
		c.isAdjacent(to, geometry.North) {
		drawBox(img, c.rect().Inset(c.size-off, off, 0, off), color)
	}
	if c.isAdjacent(from, geometry.East) ||
		c.isAdjacent(to, geometry.East) {
		drawBox(img, c.rect().Inset(off, c.size-off, off, 0), color)
	}
	if c.isAdjacent(from, geometry.South) ||
		c.isAdjacent(to, geometry.South) {
		drawBox(img, c.rect().Inset(0, off, c.size-off, off), color)
	}
	if c.isAdjacent(from, geometry.West) ||
		c.isAdjacent(to, geometry.West) {
		drawBox(img, c.rect().Inset(off, 0, off, c.size-off), color)
	}
}
