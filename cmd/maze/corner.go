package main

// cornerValue contains an identifier of a corner of a square.
type cornerValue int

const (
	topLeft cornerValue = iota
	topRight
	bottomLeft
	bottomRight
)
