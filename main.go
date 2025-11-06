package main

import (
	"fmt"

	"github.com/klahadore/shortest-path-ev/graph"
)

func main() {
	g := graph.NewGraph()

	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 8)
	g.AddEdge(1, 4, 6)
	g.AddEdge(2, 3, 2)
	g.AddEdge(4, 3, 10)

	_, prev := graph.ShortestPath(0, g)
	fmt.Println(graph.BuildPath(0, 3, prev))
}
