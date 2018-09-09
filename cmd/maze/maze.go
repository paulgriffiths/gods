package main

import (
	"github.com/paulgriffiths/gods/graphs"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// init seeds the random number generator.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// maze implements a maze.
type maze struct {
	point
	cellSize int
	cells    []byte
}

// newMaze returns an undug maze of the specified dimensions.
func newMaze(x, y, cellSize int) maze {
	m := maze{point{x, y}, cellSize, make([]byte, x*y)}
	m.generate()
	return m
}

// generate generates a random maze
func (m *maze) generate() {
	m.reset()
	m.dig(m.cell(0))
}

// reset resets a maze to its initial undug state.
func (m *maze) reset() {
	for i := range m.cells {
		m.cells[i] = northWall | eastWall | southWall | westWall | earth
	}
}

// dig recursively digs into all of the specified cell's neighbors
// which are not already dug.
func (m *maze) dig(c *cell) {
	m.setAttr(c, earth, false)
	for {
		nbs := m.earthNeighbors(c)
		if len(nbs) == 0 {
			return
		}
		i := rand.Intn(len(nbs))
		m.tunnel(c, nbs[i])
		m.dig(nbs[i])
	}
}

// tunnel removes the wall between the first cell and the second cell
// and digs the earth from the second cell.
func (m *maze) tunnel(c1, c2 *cell) {
	switch {
	case c1.isAdjacent(c2, west):
		m.setAttr(c1, eastWall, false)
		m.setAttr(c2, westWall, false)
	case c1.isAdjacent(c2, east):
		m.setAttr(c1, westWall, false)
		m.setAttr(c2, eastWall, false)
	case c1.isAdjacent(c2, north):
		m.setAttr(c1, southWall, false)
		m.setAttr(c2, northWall, false)
	case c1.isAdjacent(c2, south):
		m.setAttr(c1, northWall, false)
		m.setAttr(c2, southWall, false)
	default:
		panic("unexpected tunnel configuration")
	}
}

// setAttr turns the specified attribute on or off for the specified cell.
func (m *maze) setAttr(c *cell, attr byte, on bool) {
	if on {
		c.value |= attr
	} else {
		c.value &^= attr
	}
	m.cells[m.cellIndex(c)] = c.value
}

// neighbor returns the neighboring cell in the specified direction,
// or nil and false if there is no neighbor in that direction.
func (m maze) neighbor(c *cell, d direction) (*cell, bool) {
	p := c.point
	switch d {
	case north:
		p.y--
	case east:
		p.x++
	case south:
		p.y++
	case west:
		p.x--
	}
	if !m.inRange(p) {
		return nil, false
	}
	return m.cell(m.index(p)), true
}

// earthNeighbors returns a slice containing all the neighbors of
// the specified cell which have not yet been dug.
func (m maze) earthNeighbors(c *cell) []*cell {
	nbs := []*cell{}
	for _, dir := range []direction{north, east, south, west} {
		if nb, ok := m.neighbor(c, dir); ok {
			if nb.hasAttr(earth) {
				nbs = append(nbs, nb)
			}
		}
	}
	return nbs
}

// index converts x, y coordinates into the index of the maze's
// cell array.
func (m maze) index(p point) int {
	return p.y*m.x + p.x
}

func (m maze) indexPoint(i int) point {
	return point{i % m.x, i / m.x}
}

// cellIndex extracts from a cell the index of the maze's cell array.
func (m maze) cellIndex(c *cell) int {
	return m.index(c.point)
}

// inRange returns true if the specified coordinates are within the range
// of the maze's dimensions.
func (m maze) inRange(p point) bool {
	return p.x >= 0 && p.x < m.x && p.y >= 0 && p.y < m.y
}

// cell returns the cell at the specified coordinates.
func (m maze) cell(i int) *cell {
	if i < 0 || i >= len(m.cells) {
		panic("cell out of range")
	}
	return &cell{m.indexPoint(i), m.cellSize, m.cells[i]}
}

// drawLine draws a line of the specified length from the specified
// point in the specified direction with the specified color.
func drawLine(img *image.Paletted, p point, length int,
	dir direction, color uint8) {
	var xInc, yInc int
	switch dir {
	case north:
		xInc, yInc = 0, -1
	case east:
		xInc, yInc = 1, 0
	case south:
		xInc, yInc = 0, 1
	case west:
		xInc, yInc = -1, 0
	default:
		panic("unexpected direction")
	}
	for i := 0; i < length; i++ {
		img.SetColorIndex(p.x+xInc*i, p.y+yInc*i, color)
	}
}

// box represents a box.
type box struct {
	topLeft     point
	bottomRight point
}

// drawBox draws a solid box in the specified color.
func drawBox(img *image.Paletted, b box, color uint8) {
	for x := b.topLeft.x; x <= b.bottomRight.x; x++ {
		for y := b.topLeft.y; y <= b.bottomRight.y; y++ {
			img.SetColorIndex(x, y, color)
		}
	}
}

// insetBox returns a box derived from the provided box, with the edges
// inset by the specified amounts.
func insetBox(b box, top, right, bottom, left int) box {
	return box{
		point{b.topLeft.x + left, b.topLeft.y + top},
		point{b.bottomRight.x - right, b.bottomRight.y - bottom},
	}
}

// drawPath draws a path marker in the center of a cell.
func drawPath(img *image.Paletted, c, from, to *cell, color uint8) {
	off := c.size / 3
	if from == nil && to == nil {
		drawBox(img, insetBox(c.box(), off, off, off, off), color)
		return
	}

	if from != nil {
		if c.isAdjacent(from, north) {
			drawBox(img, insetBox(c.box(), off, off, 0, off), color)
		}
		if c.isAdjacent(from, east) {
			drawBox(img, insetBox(c.box(), off, off, off, 0), color)
		}
		if c.isAdjacent(from, south) {
			drawBox(img, insetBox(c.box(), 0, off, off, off), color)
		}
		if c.isAdjacent(from, west) {
			drawBox(img, insetBox(c.box(), off, 0, off, off), color)
		}
	}
	if to != nil {
		if c.isAdjacent(to, north) {
			drawBox(img, insetBox(c.box(), off, off, 0, off), color)
		}
		if c.isAdjacent(to, east) {
			drawBox(img, insetBox(c.box(), off, off, off, 0), color)
		}
		if c.isAdjacent(to, south) {
			drawBox(img, insetBox(c.box(), 0, off, off, off), color)
		}
		if c.isAdjacent(to, west) {
			drawBox(img, insetBox(c.box(), off, 0, off, off), color)
		}
	}
}

// image returns an image representation of the maze.
func (m maze) image(vl graphs.VertexList) image.Image {
	var palette = []color.Color{
		color.White,
		color.Black,
		color.RGBA{0xCC, 0x22, 0x22, 0xFF},
		color.RGBA{0x00, 0x77, 0x00, 0xFF},
	}
	const (
		whiteIndex = 0
		blackIndex = 1
		redIndex   = 2
		greenIndex = 3
	)
	rect := image.Rect(0, 0, m.cellSize*m.x, m.cellSize*m.y)
	img := image.NewPaletted(rect, palette)

	for i := range m.cells {
		c := m.cell(i)
		if c.hasAttr(northWall) {
			drawLine(img, c.corner(topLeft), m.cellSize,
				east, blackIndex)
		}
		if c.hasAttr(eastWall) {
			drawLine(img, c.corner(topRight), m.cellSize,
				south, blackIndex)
		}
		if c.hasAttr(southWall) {
			drawLine(img, c.corner(bottomLeft), m.cellSize,
				east, blackIndex)
		}
		if c.hasAttr(westWall) {
			drawLine(img, c.corner(topLeft), m.cellSize,
				south, blackIndex)
		}
	}

	if vl != nil {
		for n, v := range vl {
			var from, to *cell
			if n == 0 {
				from = nil
			} else {
				from = m.cell(int(vl[n-1]))
			}
			if n == len(vl)-1 {
				to = nil
			} else {
				to = m.cell(int(vl[n+1]))
			}
			c := m.cell(int(v))
			drawPath(img, c, from, to, redIndex)
		}
	}

	return img
}

// toGraph returns a graph representing the maze.
func (m maze) toGraph() graphs.Graph {
	g := graphs.NewAmGraph(len(m.cells))
	for l := range m.cells {
		c := m.cell(l)
		if !c.hasAttr(northWall) {
			r := m.index(point{c.x, c.y - 1})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
		if !c.hasAttr(eastWall) {
			r := m.index(point{c.x + 1, c.y})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
		if !c.hasAttr(southWall) {
			r := m.index(point{c.x, c.y + 1})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
		if !c.hasAttr(westWall) {
			r := m.index(point{c.x - 1, c.y})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
	}
	return g
}
