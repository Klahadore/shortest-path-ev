package main

import (
	"github.com/klahadore/shortest-path-ev/graph"
)

func main() {
	g := graph.NewGraph()

	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	g.AddEdge(4, 3)

	graph.ShortestPath(2, g)

}
