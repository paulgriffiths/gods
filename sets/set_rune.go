package sets

// SetRune implements a set of runes.
type SetRune map[rune]bool

// NewSetRune creates a new set of runes with optional initial elements.
func NewSetRune(values ...rune) SetRune {
	newSet := make(map[rune]bool)
	for _, value := range values {
		newSet[value] = true
	}
	return newSet
}

// IsEmpty returns true if a set is the empty set.
func (s SetRune) IsEmpty() bool {
	return len(s) == 0
}

// Length returns the number of elements in the set.
func (s SetRune) Length() int {
	return len(s)
}

// Elements returns an array of the elements in the set.
func (s SetRune) Elements() []rune {
	list := make([]rune, 0, len(s))
	for key := range s {
		list = append(list, key)
	}
	return list
}

// Equals tests if two sets contain the same members
func (s SetRune) Equals(other SetRune) bool {
	if len(s) != len(other) || len(s) != len(s.Union(other)) {
		return false
	}
	return true
}

// Contains returns true if the set contains the specified rune.
func (s SetRune) Contains(n rune) bool {
	return s[n]
}

// Insert inserts an rune into a set if it isn't already in the set.
func (s *SetRune) Insert(n rune) {
	(*s)[n] = true
}

// Intersection returns the intersection of two sets.
func (s SetRune) Intersection(other SetRune) SetRune {
	inter := NewSetRune()
	for key := range s {
		if other[key] {
			inter[key] = true
		}
	}
	return inter
}

// Union returns the union of two sets.
func (s SetRune) Union(other SetRune) SetRune {
	return NewSetRune(append(s.Elements(), other.Elements()...)...)
}
