package graph

const maxInt = int(^uint(0) >> 1)

func ShortestPath(source int, g *Graph) (map[int]int, map[int]int) {
	// At first it contains all the vertices, later we will remove the source because this is not necessary
	unvisited := g.GetVertices()
	prev := make(map[int]int)
	dist := make(map[int]int)

	// Initialize the map with all nodes exceptD the source to have a distance of inf because
	// they are initially unexplored.
	for node := range unvisited.elements {
		dist[node] = maxInt
		prev[node] = -1
	}
	dist[source] = 0

	for unvisited.Size() != 0 {
		// Just gets a random item from the set
		var currentNode int
		for currentNode, _ = range unvisited.elements {
			break
		}
		minDistLength := dist[currentNode]
		for node := range unvisited.elements {
			// If we haven't visited it yet, then
			if dist[node] < minDistLength {
				currentNode = node
				minDistLength = dist[node]
			}
		}

		if minDistLength == maxInt {
			break // all remaining nodes are unreachable
		}

		edges := g.adjList[currentNode]
		for _, edge := range edges {
			if unvisited.Contains(edge.To) {
				if (edge.Weight + dist[currentNode]) < dist[edge.To] {
					dist[edge.To] = edge.Weight + dist[currentNode]
					prev[edge.To] = currentNode
				}
			}
		}

		unvisited.Remove(currentNode)

	}
	return dist, prev

}

func BuildPath(source, target int, prev map[int]int) []int {
	path := []int{}

	for cur := target; cur != -1; cur = prev[cur] {
		path = append(path, cur)
		if cur == source {
			break
		}
	}

	// If we never reached the source, there's no path
	if len(path) == 0 || path[len(path)-1] != source {
		return nil // unreachable
	}

	// Reverse the path to get source -> target
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
