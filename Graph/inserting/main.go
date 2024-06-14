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

func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(3, 4)
	g.insert(1, 3)
	g.insert(1, 5)

	g.display()
}
