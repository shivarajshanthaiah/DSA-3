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

func (g *Graph) removeEdge(v1, v2 int) {
	arr := g.vertex[v1]
	for i := 0; i < len(arr); i++ {
		if arr[i] == v2 {
			g.vertex[v1] = append(arr[:1], arr[i+1:]...)
			break
		}
	}
	arr = g.vertex[v2]
	for i := 0; i < len(arr); i++ {
		if arr[i] == v1 {
			g.vertex[v2] = append(arr[:1], arr[i+1:]...)
			break
		}
	}
}

func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(2, 3)
	g.insert(4, 5)
	g.insert(5, 2)
	g.insert(1, 5)

	g.display()
	g.removeEdge(1, 5)
	g.display()
}
