package graphs

// DfsVertexList returns a VertexList containing, in order, the vertices
// visited during a recursive depth first search of a graph.
func DfsVertexList(g Graph, v Vertex) VertexList {
	l := VertexList{}
	f := func(v Vertex) {
		l = append(l, v)
	}
	Dfs(g, v, f)
	return l
}

// DfsVertexListIterative returns a VertexList containing, in order,
// the vertices visited during an iterative depth first search of a graph.
func DfsVertexListIterative(g Graph, v Vertex) VertexList {
	l := VertexList{}
	f := func(v Vertex) {
		l = append(l, v)
	}
	DfsIterative(g, v, f)
	return l
}

// Dfs performs a recursive depth first search of a graph.
func Dfs(g Graph, v Vertex, f func(Vertex)) {
	g.Visit(v)
	f(v)
	for _, x := range g.Adjacent(v) {
		if !g.IsVisited(x) {
			Dfs(g, x, f)
		}
	}
}

// DfsIterative performs an iterative depth first search of a graph.
func DfsIterative(g Graph, v Vertex, f func(Vertex)) {
	s := NewStackVertex()
	s.Push(v)
	for !s.IsEmpty() {
		vtx := s.Pop()
		if !g.IsVisited(vtx) {
			g.Visit(vtx)
			f(vtx)
			a := g.Adjacent(vtx)
			for i := len(a) - 1; i >= 0; i-- {
				s.Push(a[i])
			}
		}
	}
}

// BfsVertexList returns a VertexList containing, in order, the vertices
// visited during a breadth first search of a graph.
func BfsVertexList(g Graph, v Vertex) VertexList {
	l := VertexList{}
	f := func(v Vertex) {
		l = append(l, v)
	}
	Bfs(g, v, f)
	return l
}

// Bfs performs a breadth first search of a graph.
func Bfs(g Graph, v Vertex, f func(Vertex)) {
	q := NewQueueVertex()
	g.Visit(v)
	f(v)
	q.Enqueue(v)
	for !q.IsEmpty() {
		vtx := q.Dequeue()
		a := g.Adjacent(vtx)
		for i := 0; i < len(a); i++ {
			if !g.IsVisited(a[i]) {
				g.Visit(a[i])
				f(a[i])
				q.Enqueue(a[i])
			}
		}
	}
}

// BfsFindPath finds, via a breadth first search, the a path between
// vertices start and target, and returns a VertexList representing
// that path.
func BfsFindPath(g Graph, start, target Vertex) VertexList {
	pred := make([]Vertex, g.NumVertices())
	for i := 0; i < g.NumVertices(); i++ {
		pred[i] = Vertex(-1)
	}
	found := false

	q := NewQueueVertex()
	g.Visit(start)
	q.Enqueue(start)

outer:
	for !q.IsEmpty() {
		vtx := q.Dequeue()
		a := g.Adjacent(vtx)
		for i := 0; i < len(a); i++ {
			if !g.IsVisited(a[i]) {
				pred[a[i]] = vtx
				g.Visit(a[i])
				q.Enqueue(a[i])
				if a[i] == target {
					found = true
					break outer
				}
			}
		}
	}

	if !found {
		return nil
	}

	path := VertexList{}
	crawl := target
	path = append(path, crawl)
	for pred[crawl] != -1 {
		path = append(path, pred[crawl])
		crawl = pred[crawl]
	}
	return path
}
