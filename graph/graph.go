package graph

import (
	"fmt"
	"sort"
)

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	adjList map[int][]Edge
}

// The function is a constructor for the Graph struct
// It returns a pointer to a Graph object.
func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]Edge),
	}
}

func (g *Graph) AddVertex(vertex int) {
	_, exists := g.adjList[vertex]
	if !exists {
		g.adjList[vertex] = []Edge{}
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)

	g.adjList[from] = append(g.adjList[from], Edge{To: to, Weight: weight})
}

func (g *Graph) GetVertices() []int {
	vertices := make([]int, 0, len(g.adjList))
	for vertex := range g.adjList {
		vertices = append(vertices, vertex)
	}
	sort.Ints(vertices)
	return vertices
}

func (g *Graph) Print() {
	for vertex, neighbors := range g.adjList {
		// Print each vertex followed by its neighbor list
		// Use integer and slice formatting with a newline per entry
		fmt.Printf("%d -> %v\n", vertex, neighbors)
	}
}
