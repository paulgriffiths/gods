package graphs

// Graph represents a graph type.
type Graph interface {
	InsertEdge(v1, v2 Vertex)
	Adjacent(v Vertex) VertexList
	Visit(v Vertex)
	IsVisited(v Vertex) bool
	NumVertices() int
}
