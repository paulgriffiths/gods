package graphs

type alGraph struct {
	matrix  [][]Vertex
	visited []bool
	size    int
}

// NewAlGraph returns a new graph with vertices 0...n-1 and no edges,
// implemented as an adjacency matrix.
func NewAlGraph(n int) Graph {
	m := make([][]Vertex, n)
	for i := range m {
		m[i] = make([]Vertex, 0, 8) // Assume reasonable default capacity
	}
	v := make([]bool, n)
	return &alGraph{m, v, n}
}

func insertIfNotPresent(s *[]Vertex, v Vertex) {
	for _, i := range *s {
		if i == v {
			return
		}
	}
	*s = append(*s, v)
}

func (g *alGraph) InsertEdge(v1, v2 Vertex) {
	insertIfNotPresent(&g.matrix[v1], v2)
	insertIfNotPresent(&g.matrix[v2], v1)
}

func (g *alGraph) Adjacent(v Vertex) VertexList {
	return g.matrix[v]
}

func (g *alGraph) Visit(v Vertex) {
	g.visited[v] = true
}

func (g *alGraph) IsVisited(v Vertex) bool {
	return g.visited[v] == true
}

func (g *alGraph) NumVertices() int {
	return g.size
}
