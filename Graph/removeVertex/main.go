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

func (g *Graph) removeVertex(ver int) {
	arr := g.vertex[ver]
	for i := 0; i < len(arr); i++ {
		edges := g.vertex[arr[i]]
		for j := 0; j < len(arr); j++ {
			if edges[j] == ver {
				g.vertex[arr[i]] = append(edges[:j], edges[j+1:]...)
				break
			}
		}
	}
	delete(g.vertex, ver)

}

func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(2, 3)
	g.insert(4, 5)
	g.insert(5, 2)
	g.insert(3, 5)

	g.display()
	g.removeVertex(1)
	g.display()
}
