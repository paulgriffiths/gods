package main

import (
	"flag"
	"fmt"
	"github.com/paulgriffiths/gods/graphs"
	"image/png"
	"os"
)

func main() {
	x := flag.Int("x", 25, "x dimension of maze")
	y := flag.Int("y", 25, "y dimension of maze")
	c := flag.Int("s", 30, "dimension of each maze cell in pixels")
	f := flag.String("o", "stdout", "output file")
	p := flag.Bool("p", false, "show path through maze")
	flag.Parse()

	m := newMaze(*x, *y, *c)
	g := m.toGraph()

	var l graphs.VertexList
	if *p {
		l = graphs.BfsFindPath(g, 0, graphs.Vertex(*x**y-1))
	} else {
		l = nil
	}

	img := m.image(l)

	if *f == "stdout" {
		png.Encode(os.Stdout, img)
	} else {
		outfile, err := os.Create(*f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't open output file: %v\n", err)
			os.Exit(1)
		}
		defer outfile.Close()
		png.Encode(outfile, img)
	}
}
