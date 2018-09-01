package graphs

import (
	"sort"
)

// EdgeList represents a list of undirected graph edges.
type EdgeList []Edge

// Len returns the length of an EdgeList
func (e EdgeList) Len() int {
	return len(e)
}

// Equals returns true if the EdgeList is equal to the provided EdgeList.
func (e EdgeList) Equals(other EdgeList) bool {
	if e.Len() != other.Len() {
		return false
	}
	for n := range e {
		if !e[n].Equals(other[n]) {
			return false
		}
	}
	return true
}

// Less returns true if the edge at index i is less than the edge
// at index j.
func (e EdgeList) Less(i, j int) bool {
	return e[i].Less(e[j])
}

// Swap swaps the edges at indices i and j.
func (e EdgeList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Sort sorts an EdgeList in-place.
func (e EdgeList) Sort() {
	sort.Sort(e)
}
