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

func (g *Graph) hasPath(a, b int) (bool, int) {
	visited := make(map[int]bool)
	count := 0
	arr := []int{a}

	for len(arr) > 0 {
		val := arr[0]
		arr = arr[1:]
		visited[val] = true
		count++
		for _, v := range g.vertex[val] {
			if v == b {
				return true, count
			}
			if !visited[v] {
				arr = append(arr, v)
				visited[v] = true
			}
		}
	}
	return false, 0
}
func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(4, 5)
	g.insert(3, 2)
	g.insert(4, 6)
	g.insert(3, 7)
	g.insert(3, 5)
	g.display()
	fmt.Println(g.hasPath(1,6))

}
