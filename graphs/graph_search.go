package graphs

import (
	"github.com/paulgriffiths/gods/queues"
	"github.com/paulgriffiths/gods/stacks"
)

// DfsSpan returns a depth first spanning tree for a graph.
func DfsSpan(g Graph, v int) []int {
	l := []int{}
	f := func(n int) {
		l = append(l, n)
	}
	Dfs(g, v, f)
	return l
}

// DfsSpanIterative returns a depth first spanning tree for a graph.
func DfsSpanIterative(g Graph, v int) []int {
	l := []int{}
	f := func(n int) {
		l = append(l, n)
	}
	DfsIterative(g, v, f)
	return l
}

// Dfs performs a recursive depth first search of a graph.
func Dfs(g Graph, v int, f func(int)) {
	g.Visit(v)
	f(v)
	for _, x := range g.Adjacent(v) {
		if !g.IsVisited(x) {
			Dfs(g, x, f)
		}
	}
}

// DfsIterative performs an iterative depth first search of a graph.
func DfsIterative(g Graph, v int, f func(int)) {
	s := stacks.NewStackInt()
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

// BfsSpan returns a breadth first spanning tree for a graph.
func BfsSpan(g Graph, v int) []int {
	l := []int{}
	f := func(n int) {
		l = append(l, n)
	}
	Bfs(g, v, f)
	return l
}

// Bfs performs a breadth first search of a graph.
func Bfs(g Graph, v int, f func(int)) {
	q := queues.NewQueueInt()
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
