package sets

// SetInt implements a set of integers
type SetInt []int

// IsEmpty returns true if a set is the empty set.
func (s SetInt) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// EmptySet returns the empty set
func EmptySet() SetInt {
	return SetInt{}
}

// Equals tests if two sets contain the same members
func (s SetInt) Equals(other SetInt) bool {
	if len(s) != len(other) {
		return false
	}
	for _, element := range s {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// Contains returns true if the set contains the specified integer.
func (s SetInt) Contains(n int) bool {
	for _, element := range s {
		if element == n {
			return true
		}
	}
	return false
}

// Insert inserts an integer into a set if it isn't already in the set.
func (s *SetInt) Insert(n int) {
	if !s.Contains(n) {
		*s = append(*s, n)
	}
}

// Intersection returns the intersection of two sets.
func (s SetInt) Intersection(other SetInt) SetInt {
	newSet := SetInt{}
	for _, element := range s {
		if other.Contains(element) {
			newSet.Insert(element)
		}
	}
	return newSet
}

// Union returns the union of two sets.
func (s SetInt) Union(other SetInt) SetInt {
	newSet := SetInt{}
	for _, element := range s {
		newSet.Insert(element)
	}
	for _, element := range other {
		newSet.Insert(element)
	}
	return newSet
}
