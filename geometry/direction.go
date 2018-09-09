package geometry

// Direction represents a cardinal direction.
type Direction uint8

const (
	// North represents the north compass direction.
	North Direction = iota
	// East represents the east compass direction.
	East
	// South represents the south compass direction.
	South
	// West represents the west compass direction.
	West
)

// Directions is an array containing all the cardinal directions.
var Directions = [4]Direction{North, East, South, West}
