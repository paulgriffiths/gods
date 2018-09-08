package sets

// SetInt implements a set of integers
type SetInt []int

// Contains returns true if the set contains the specified integer.
func (s *SetInt) Contains(n int) bool {
	for _, element := range *s {
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
