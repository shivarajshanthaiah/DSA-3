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

func (g *Graph) bfs(val int) {
	visited := make(map[int]bool)
	visited[val] = true
	arr := []int{val}
	for len(arr) > 0 {
		v := arr[0]
		fmt.Printf("%d", v)
		arr = arr[1:]

		for _, e := range g.vertex[v] {
			if !visited[e] {
				arr = append(arr, e)
				visited[e] = true
			}
		}
	}
}

func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(2, 3)
	g.insert(5, 3)
	g.insert(5, 7)
	g.insert(7, 9)
	g.bfs(9)
}
