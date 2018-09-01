package graphs

type amGraph struct {
	matrix  [][]byte
	visited []bool
	size    int
}

// NewAmGraph returns a new graph with vertices 0...n-1 and no edges,
// implemented as an adjacency matrix.
func NewAmGraph(n int) Graph {
	m := make([][]byte, n)
	for i := range m {
		m[i] = make([]byte, n)
	}
	v := make([]bool, n)
	return &amGraph{m, v, n}
}

func (g *amGraph) InsertEdge(v1, v2 Vertex) {
	g.matrix[v1][v2] |= 1
	g.matrix[v2][v1] |= 1
}

func (g *amGraph) Adjacent(v Vertex) VertexList {
	result := VertexList{}
	for n := 0; n < len(g.matrix[v]); n++ {
		if g.matrix[v][n]&0x01 == 0x01 {
			result = append(result, Vertex(n))
		}
	}
	return result
}

func (g *amGraph) Visit(v Vertex) {
	g.visited[v] = true
}

func (g *amGraph) IsVisited(v Vertex) bool {
	return g.visited[v] == true
}

func (g *amGraph) NumVertices() int {
	return g.size
}
