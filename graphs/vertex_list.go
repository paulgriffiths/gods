package graphs

import (
	"sort"
)

// VertexList represents a list of graph vertices.
type VertexList []Vertex

// Len returns the length of a VertexList.
func (v VertexList) Len() int {
	return len(v)
}

// Equals returns true if the VertexList is equal to the provided
// VertexList.
func (v VertexList) Equals(other VertexList) bool {
	if v.Len() != other.Len() {
		return false
	}
	for n := range v {
		if !v[n].Equals(other[n]) {
			return false
		}
	}
	return true
}

// Less returns true if the vertex at index i is less than the vertex
// at index j.
func (v VertexList) Less(i, j int) bool {
	return v[i].Less(v[j])
}

// Swap swaps the vertices at indices i and j
func (v VertexList) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// Sort sorts a VertexList in-place.
func (v VertexList) Sort() {
	sort.Sort(v)
}
