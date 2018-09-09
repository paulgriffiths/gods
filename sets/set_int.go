package sets

// SetInt implements a set of integers
type SetInt struct {
	set SetInterface
}

// NewSetInt creates a new set of integers.
func NewSetInt(values ...int) SetInt {
	newSet := NewSetInterface(func(a, b interface{}) bool {
		return a == b
	})
	for _, value := range values {
		newSet.Insert(value)
	}
	return SetInt{newSet}
}

// IsEmpty returns true if a set is the empty set.
func (s SetInt) IsEmpty() bool {
	return s.set.IsEmpty()
}

// Length returns the number of elements in the set.
func (s SetInt) Length() int {
	return s.set.Length()
}

// Elements returns an array of the elements in the set.
func (s SetInt) Elements() []int {
	list := []int{}
	for _, e := range s.set.Elements() {
		list = append(list, e.(int))
	}
	return list
}

// Equals tests if two sets contain the same members
func (s SetInt) Equals(other SetInt) bool {
	return s.set.Equals(other.set)
}

// Contains returns true if the set contains the specified integer.
func (s SetInt) Contains(n int) bool {
	return s.set.Contains(n)
}

// Insert inserts an integer into a set if it isn't already in the set.
func (s *SetInt) Insert(n int) {
	s.set.Insert(n)
}

// Intersection returns the intersection of two sets.
func (s SetInt) Intersection(other SetInt) SetInt {
	return SetInt{s.set.Intersection(other.set)}
}

// Union returns the union of two sets.
func (s SetInt) Union(other SetInt) SetInt {
	return SetInt{s.set.Union(other.set)}
}
