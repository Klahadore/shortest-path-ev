package graph

// import (
// 	"fmt"
// 	"slices"
// )

// const maxInt = int(^uint(0) >> 1)

// // Source - https://stackoverflow.com/questions/62369223/the-equivalent-of-indexof
// // Posted by thwd
// // Retrieved 11/4/2025, License - CC-BY-SA 4.0

// // IndexOf returns the first index of needle in haystack
// // or -1 if needle is not in haystack.
// func indexOf(haystack []int, needle int) int {
// 	for i, v := range haystack {
// 		if v == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func ShortestPath(source int, g *Graph) {
// 	// At first it contains all the vertices, later we will remove the source because this is not necessary
// 	unvisited := g.GetVertices()
// 	shortestPathLength := make(map[int]int)

// 	// This just initializes the map.
// 	for node := range unvisited {
// 		if node == source {
// 			shortestPathLength[node] = 0
// 			continue
// 		}
// 		shortestPathLength[node] = maxInt
// 	}

// 	indexOfSource := indexOf(unvisited, source)
// 	unvisited = slices.Delete(unvisited, indexOfSource, indexOfSource+1)

// 	currentNode := source
// 	for len(unvisited) != 0 {
// 		neighbors := g.adjList[currentNode]

// 		minLengthNeighbor := neighbors[0]
// 		for neighbor := range neighbors {
// 			if (shortestPathLength[current] + neighbor.Weight) < shortestpathLength[neighbor] {

// 			}
// 		}
// 	}

// 	fmt.Println(unvisited)
// 	fmt.Println(shortestPathLength)

// }
