package main

import "fmt"

type Graph struct {
	vertex map[int][]int
}

func newGraph() *Graph {
	return &Graph{vertex: make(map[int][]int)}
}

func (g *Graph) insert(v1, v2 int) {
	g.vertex[v1] = append(g.vertex[v1], v2)
	g.vertex[v2] = append(g.vertex[v2], v1)
}

func (g *Graph) dfs(val int, visited map[int]bool) {
	visited[val] = true
	for _, neighbor := range g.vertex[val] {
		if !visited[neighbor] {
			g.dfs(neighbor, visited)
		}
	}
}

func (g *Graph) canVisitAll() bool {
	visited := make(map[int]bool)
	for v := range g.vertex {
		visited[v] = false
	}

	var startVertex int
	for v := range g.vertex {
		startVertex = v
		break
	}

	g.dfs(startVertex, visited)

	for _, visit := range visited {
		if !visit {
			return false
		}
	}
	return true
}

func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(4, 5)
	g.insert(3, 2)
	g.insert(4, 6)
	g.insert(3, 7)
	g.insert(3, 5)

	fmt.Println(g.canVisitAll())
}
