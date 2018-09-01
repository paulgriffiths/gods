package amGraph

type amGraph struct {
	matrix  [][]byte
	visited []bool
	size    int
}

func NewAmGraph(n int) *amGraph {
	m := make([][]byte, n)
	for i, _ := range m {
		m[i] = make([]byte, n)
	}
	v := make([]bool, n)
	return &amGraph{m, v, n}
}

func (g *amGraph) InsertEdge(v1, v2 int) {
	g.matrix[v1][v2] |= 1
	g.matrix[v2][v1] |= 1
}

func (g *amGraph) Adjacent(v int) []int {
	result := []int{}
	for n := 0; n < len(g.matrix[v]); n++ {
		if g.matrix[v][n]&0x01 == 0x01 {
			result = append(result, n)
		}
	}
	return result
}

func (g *amGraph) Visit(v int) {
	g.visited[v] = true
}

func (g *amGraph) IsVisited(v int) bool {
	return g.visited[v] == true
}

func (g *amGraph) NumVertices() int {
	return g.size
}
