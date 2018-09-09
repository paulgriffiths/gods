package main

import (
	"flag"
	"fmt"
	"github.com/paulgriffiths/gods/graphs"
	"image/png"
	"os"
)

func main() {
	x := flag.Int("x", 25, "x dimension of maze, >= 2")
	y := flag.Int("y", 25, "y dimension of maze, >= 2")
	c := flag.Int("s", 30, "dimension of each maze cell in pixels, >= 6")
	f := flag.String("o", "stdout", "output file")
	p := flag.Bool("p", false, "show path through maze")
	flag.Parse()

	argError := false
	if *x < 2 {
		fmt.Fprintf(os.Stderr, "%s: x dimension must be at least 2\n",
			os.Args[0])
		argError = true
	}
	if *y < 2 {
		fmt.Fprintf(os.Stderr, "%s: y dimension must be at least 2\n",
			os.Args[0])
		argError = true
	}
	if *c < 6 {
		fmt.Fprintf(os.Stderr, "%s: cell dimension must be at least 6\n",
			os.Args[0])
		argError = true
	}

	if argError {
		os.Exit(1)
	}

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
