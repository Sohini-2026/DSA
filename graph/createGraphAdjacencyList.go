package graph

func CreateGraphAdjacencyList(edges [][]int) map[int][]int {
	graph := make(map[int][]int)

	for _, edge := range edges {
		if len(edge) != 2 {
			continue // Skip invalid edges
		}
		u, v := edge[0], edge[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // For undirected graph
	}

	return graph
}

// Example usage:
// edges := [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}}
// graph := CreateGraphAdjacencyList(edges)
