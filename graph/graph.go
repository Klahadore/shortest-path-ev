package graph

import (
	"fmt"
)

type Graph struct {
	adjList map[int][]int
}

// The function is a constructor for the Graph struct
// It returns a pointer to a Graph object.
func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func (g *Graph) AddVertex(vertex int) {
	_, exists := g.adjList[vertex]
	if !exists {
		g.adjList[vertex] = []int{}
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.AddVertex(from)
	g.AddVertex(to)

	g.adjList[from] = append(g.adjList[from], to)
}

func (g *Graph) Print() {
    for vertex, neighbors := range g.adjList {
        // Print each vertex followed by its neighbor list
        // Use integer and slice formatting with a newline per entry
        fmt.Printf("%d -> %v\n", vertex, neighbors)
    }
}
