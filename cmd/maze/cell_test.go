package main

import (
	"github.com/paulgriffiths/gods/geometry"
	"testing"
)

var m maze

/*
 Representation of test maze:

 =======
 | |X| |
 |___  |
 |X|X|_|
 =======
*/

var b = []byte{
	northWall | eastWall | westWall,
	westWall | northWall | eastWall | southWall | earth,
	northWall | eastWall | westWall,
	westWall | southWall,
	southWall,
	eastWall,
	westWall | northWall | eastWall | southWall | earth,
	westWall | northWall | eastWall | southWall | earth,
	westWall | eastWall | southWall,
}

func init() {
	m = newMaze(3, 3, 30)
	m.cells = b
}

func TestCellValues(t *testing.T) {
	inputs := []struct {
		x, y  int
		value byte
	}{
		{0, 0, b[0]},
		{1, 0, b[1]},
		{2, 0, b[2]},
		{0, 1, b[3]},
		{1, 1, b[4]},
		{2, 1, b[5]},
		{0, 2, b[6]},
		{1, 2, b[7]},
		{2, 2, b[8]},
	}

	for _, i := range inputs {
		c := m.cell(m.index(geometry.Point{i.x, i.y}))
		if c.value != i.value {
			t.Errorf("got %v, want %v", c.value, i.value)
		}
	}
}

func TestCellEarthNeighbors(t *testing.T) {
	inputs := []struct {
		x, y int
		vals []byte
	}{
		{0, 0, []byte{b[1]}},
		{1, 0, []byte{}},
		{2, 0, []byte{b[1]}},
		{0, 1, []byte{b[6]}},
		{1, 1, []byte{b[1], b[7]}},
		{2, 1, []byte{}},
		{0, 2, []byte{b[7]}},
		{1, 2, []byte{b[6]}},
		{2, 2, []byte{b[7]}},
	}
	for _, i := range inputs {
		c := m.cell(m.index(geometry.Point{i.x, i.y}))
		n := m.earthNeighbors(c)
		for j := range n {
			if n[j].value != i.vals[j] {
				t.Errorf("got %v, want %v", n[j].value, i.vals[j])
			}
		}
	}
}
