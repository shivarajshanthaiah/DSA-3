package main

import (
	"fmt"
	"math"
)

type Grapgh struct {
	vertex map[int][]int
}

func newGraph() *Grapgh {
	return &Grapgh{vertex: make(map[int][]int)}
}

func (g *Grapgh) insert(v1, v2 int) {
	g.vertex[v1] = append(g.vertex[v1], v2)
	g.vertex[v2] = append(g.vertex[v2], v1)
}

func (g *Grapgh) display() {
	fmt.Println(g.vertex)
}

func (g *Grapgh) findMin() int {
	minval := math.MaxInt64
	for val := range g.vertex {
		if val < minval {
			minval = val
		}
	}
	return minval
}

func main() {
	g := newGraph()
	g.insert(1, 2)
	g.insert(3, 4)
	g.insert(1, 3)
	g.insert(1, 5)

	g.display()
	fmt.Println(g.findMin())
}
