package sets

// SetRune implements a set of runes.
type SetRune struct {
	set SetInterface
}

// NewSetRune creates a new set of runes with optional initial elements.
func NewSetRune(values ...rune) SetRune {
	newSet := NewSetInterface(func(a, b interface{}) bool {
		return a.(rune) == b.(rune)
	})
	for _, value := range values {
		newSet.Insert(value)
	}
	return SetRune{newSet}
}

// IsEmpty returns true if a set is the empty set.
func (s SetRune) IsEmpty() bool {
	return s.set.IsEmpty()
}

// Length returns the number of elements in the set.
func (s SetRune) Length() int {
	return s.set.Length()
}

// Elements returns an array of the elements in the set.
func (s SetRune) Elements() []rune {
	list := []rune{}
	for _, elem := range s.set.Elements() {
		list = append(list, elem.(rune))
	}
	return list
}

// Equals tests if two sets contain the same members
func (s SetRune) Equals(other SetRune) bool {
	return s.set.Equals(other.set)
}

// Contains returns true if the set contains the specified rune.
func (s SetRune) Contains(n rune) bool {
	return s.set.Contains(n)
}

// Insert inserts an rune into a set if it isn't already in the set.
func (s *SetRune) Insert(n rune) {
	s.set.Insert(n)
}

// Intersection returns the intersection of two sets.
func (s SetRune) Intersection(other SetRune) SetRune {
	return SetRune{s.set.Intersection(other.set)}
}

// Union returns the union of two sets.
func (s SetRune) Union(other SetRune) SetRune {
	return SetRune{s.set.Union(other.set)}
}
