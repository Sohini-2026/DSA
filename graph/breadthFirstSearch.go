package graph

func BreadthFirstTraversal(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	result := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// Example usage:
// graph := CreateGraphAdjacencyList([][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}})
// bfsResult := BFS(graph, 1)
// fmt.Println(bfsResult) // Output: [1 2 3 4] or similar order depending on the graph structure
// Note: The order of traversal may vary based on the graph structure and the order of neighbors in the adjacency list.
// This BFS implementation assumes an undirected graph. For directed graphs, you can remove the line that appends the reverse edge.
