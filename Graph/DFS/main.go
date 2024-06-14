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

func (g *Graph) display() {
	fmt.Println(g.vertex)
}

func (g *Graph) dfs(val int, visited map[int]bool) {
	visited[val] = true
	fmt.Printf("%d", val)
	for _, v := range g.vertex[val] {
		if !visited[v] {
			g.dfs(v, visited)
		}
	}
}
func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(4, 5)
	g.insert(3, 2)
	g.insert(4, 6)
	g.insert(3, 7)
	g.insert(3,5)
	g.display()
	visited := make(map[int]bool)
	g.dfs(1,visited)
}
