package lib

import "testing"

func TestGraoh(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	bfs := []interface{}{1, 2, 3, 4}
	for i, v := range g.BFS(1) {
		if v != bfs[i] {
			t.Fatalf("graph BFS(1) = %v", g.BFS(1))
		}
	}
}
