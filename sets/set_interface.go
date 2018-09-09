package sets

// SetInterface implements a set of interface{} elements.
type SetInterface struct {
	elems  []interface{}               // Set elements
	eqFunc func(a, b interface{}) bool // Element equality function
}

// NewSetInterface returns a new interface{} set with the
// provided element equality function and optional initial
// list of elements. The element equality function should
// return true if elements a and b are equal.
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
	return s.Length() == 0
}

// Length returns the number of elements in the set.
func (s SetInterface) Length() int {
	return len(s.elems)
}

// Elements returns an array of the elements in the set.
func (s SetInterface) Elements() []interface{} {
	list := []interface{}{}
	for _, elem := range s.elems {
		list = append(list, elem)
	}
	return list
}

// Equals tests if two sets contain the same elements
func (s SetInterface) Equals(other SetInterface) bool {
	if s.Length() != other.Length() {
		return false
	}
	for _, elem := range s.elems {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Contains returns true if the set contains the specified element.
func (s SetInterface) Contains(e interface{}) bool {
	for _, elem := range s.elems {
		if s.eqFunc(elem, e) {
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
	for _, elem := range s.elems {
		if other.Contains(elem) {
			newSet.Insert(elem)
		}
	}
	return newSet
}

// Union returns the union of two sets.
func (s SetInterface) Union(other SetInterface) SetInterface {
	newSet := NewSetInterface(s.eqFunc)
	for _, elem := range s.elems {
		newSet.Insert(elem)
	}
	for _, elem := range other.elems {
		newSet.Insert(elem)
	}
	return newSet
}
