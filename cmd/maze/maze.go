package main

import (
	"github.com/paulgriffiths/gods/geometry"
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
	geometry.Point
	cellSize int
	cells    []byte
}

// newMaze returns an undug maze of the specified dimensions.
func newMaze(x, y, cellSize int) maze {
	m := maze{geometry.Point{x, y}, cellSize, make([]byte, x*y)}
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
	case c1.isAdjacent(c2, geometry.West):
		m.setAttr(c1, eastWall, false)
		m.setAttr(c2, westWall, false)
	case c1.isAdjacent(c2, geometry.East):
		m.setAttr(c1, westWall, false)
		m.setAttr(c2, eastWall, false)
	case c1.isAdjacent(c2, geometry.North):
		m.setAttr(c1, southWall, false)
		m.setAttr(c2, northWall, false)
	case c1.isAdjacent(c2, geometry.South):
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
func (m maze) neighbor(c *cell, d geometry.Direction) (*cell, bool) {
	p := c.Point
	switch d {
	case geometry.North:
		p.Y--
	case geometry.East:
		p.X++
	case geometry.South:
		p.Y++
	case geometry.West:
		p.X--
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
	for _, dir := range geometry.Directions {
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
func (m maze) index(p geometry.Point) int {
	return p.Y*m.X + p.X
}

func (m maze) indexPoint(i int) geometry.Point {
	return geometry.Point{i % m.X, i / m.X}
}

// cellIndex extracts from a cell the index of the maze's cell array.
func (m maze) cellIndex(c *cell) int {
	return m.index(c.Point)
}

// inRange returns true if the specified coordinates are within the range
// of the maze's dimensions.
func (m maze) inRange(p geometry.Point) bool {
	return p.X >= 0 && p.X < m.X && p.Y >= 0 && p.Y < m.Y
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
func drawLine(img *image.Paletted, p geometry.Point, length int,
	dir geometry.Direction, color uint8) {
	var xInc, yInc int
	switch dir {
	case geometry.North:
		xInc, yInc = 0, -1
	case geometry.East:
		xInc, yInc = 1, 0
	case geometry.South:
		xInc, yInc = 0, 1
	case geometry.West:
		xInc, yInc = -1, 0
	default:
		panic("unexpected direction")
	}
	for i := 0; i < length; i++ {
		img.SetColorIndex(p.X+xInc*i, p.Y+yInc*i, color)
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
	if from == nil && to == nil {
		drawBox(img, c.box().Inset(off, off, off, off), color)
		return
	}

	if from != nil {
		if c.isAdjacent(from, geometry.North) {
			drawBox(img, c.box().Inset(off, off, 0, off), color)
		}
		if c.isAdjacent(from, geometry.East) {
			drawBox(img, c.box().Inset(off, off, off, 0), color)
		}
		if c.isAdjacent(from, geometry.South) {
			drawBox(img, c.box().Inset(0, off, off, off), color)
		}
		if c.isAdjacent(from, geometry.West) {
			drawBox(img, c.box().Inset(off, 0, off, off), color)
		}
	}
	if to != nil {
		if c.isAdjacent(to, geometry.North) {
			drawBox(img, c.box().Inset(off, off, 0, off), color)
		}
		if c.isAdjacent(to, geometry.East) {
			drawBox(img, c.box().Inset(off, off, off, 0), color)
		}
		if c.isAdjacent(to, geometry.South) {
			drawBox(img, c.box().Inset(0, off, off, off), color)
		}
		if c.isAdjacent(to, geometry.West) {
			drawBox(img, c.box().Inset(off, 0, off, off), color)
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
	rect := image.Rect(0, 0, m.cellSize*m.X, m.cellSize*m.Y)
	img := image.NewPaletted(rect, palette)

	for i := range m.cells {
		c := m.cell(i)
		if c.hasAttr(northWall) {
			drawLine(img, c.corner(geometry.TopLeft), m.cellSize,
				geometry.East, blackIndex)
		}
		if c.hasAttr(eastWall) {
			drawLine(img, c.corner(geometry.TopRight), m.cellSize,
				geometry.South, blackIndex)
		}
		if c.hasAttr(southWall) {
			drawLine(img, c.corner(geometry.BottomLeft), m.cellSize,
				geometry.East, blackIndex)
		}
		if c.hasAttr(westWall) {
			drawLine(img, c.corner(geometry.TopLeft), m.cellSize,
				geometry.South, blackIndex)
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
			r := m.index(geometry.Point{c.X, c.Y - 1})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
		if !c.hasAttr(eastWall) {
			r := m.index(geometry.Point{c.X + 1, c.Y})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
		if !c.hasAttr(southWall) {
			r := m.index(geometry.Point{c.X, c.Y + 1})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
		if !c.hasAttr(westWall) {
			r := m.index(geometry.Point{c.X - 1, c.Y})
			g.InsertEdge(graphs.Vertex(l), graphs.Vertex(r))
		}
	}
	return g
}
