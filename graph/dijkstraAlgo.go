package graph

func DijkstraAlgoWithPath(adjList map[int][][]int, start int, target int) ([]int, int) {
	distances := make(map[int]int)
	visited := make(map[int]bool)
	prev := make(map[int]int)

	for node := range adjList {
		distances[node] = 1<<31 - 1 // Initialize distances to infinity
	}
	distances[start] = 0

	for len(visited) < len(adjList) {
		minNode := -1
		minDist := 1<<31 - 1
		for node, dist := range distances {
			if !visited[node] && dist < minDist {
				minDist = dist
				minNode = node
			}
		}

		if minNode == -1 {
			break
		}

		visited[minNode] = true

		for _, neighbor := range adjList[minNode] {
			nextNode, weight := neighbor[0], neighbor[1]
			if !visited[nextNode] {
				newDist := distances[minNode] + weight
				if newDist < distances[nextNode] {
					distances[nextNode] = newDist
					prev[nextNode] = minNode
				}
			}
		}
	}

	// Reconstruct path from start to target
	path := []int{}
	for at := target; at != start; at = prev[at] {
		path = append([]int{at}, path...)
		if _, ok := prev[at]; !ok {
			return nil, -1 // No path found
		}
	}
	path = append([]int{start}, path...)

	return path, distances[target]
}

func CreateGraphAdjacencyListWithWeights(edges [][]int) map[int][][]int {
	adjList := make(map[int][][]int)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		adjList[u] = append(adjList[u], []int{v, w})
		adjList[v] = append(adjList[v], []int{u, w}) // For undirected graph
	}
	return adjList
}
