package sets

// SetInterface implements a set of interface{} elements.
type SetInterface struct {
	elems  []interface{}
	eqFunc func(a, b interface{}) bool
}

// NewSetInterface returns a new interface{} set with the
// provided equality function.
func NewSetInterface(eqFunc func(a, b interface{}) bool,
	values ...interface{}) SetInterface {
	newSet := SetInterface{[]interface{}{}, eqFunc}
	for _, value := range values {
		newSet.Insert(value)
	}
	return newSet
}

// IsEmpty returns true if a set is the empty set.
func (s SetInterface) IsEmpty() bool {
	if len(s.elems) == 0 {
		return true
	}
	return false
}

// Length returns the number of elements in the set.
func (s SetInterface) Length() int {
	return len(s.elems)
}

// Elements returns an array of the elements in the set.
func (s SetInterface) Elements() []interface{} {
	list := []interface{}{}
	for _, e := range s.elems {
		list = append(list, e)
	}
	return list
}

// Equals tests if two sets contain the same elements
func (s SetInterface) Equals(other SetInterface) bool {
	if s.Length() != other.Length() {
		return false
	}
	for _, element := range s.elems {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// Contains returns true if the set contains the specified element.
func (s SetInterface) Contains(e interface{}) bool {
	for _, element := range s.elems {
		if s.eqFunc(element, e) {
			return true
		}
	}
	return false
}

// Insert inserts an element into a set if it isn't already in the set.
func (s *SetInterface) Insert(e interface{}) {
	if !s.Contains(e) {
		s.elems = append(s.elems, e)
	}
}

// Intersection returns the intersection of two sets.
func (s SetInterface) Intersection(other SetInterface) SetInterface {
	newSet := NewSetInterface(s.eqFunc)
	for _, element := range s.elems {
		if other.Contains(element) {
			newSet.Insert(element)
		}
	}
	return newSet
}

// Union returns the union of two sets.
func (s SetInterface) Union(other SetInterface) SetInterface {
	newSet := NewSetInterface(s.eqFunc)
	for _, element := range s.elems {
		newSet.Insert(element)
	}
	for _, element := range other.elems {
		newSet.Insert(element)
	}
	return newSet
}
