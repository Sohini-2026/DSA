package graph

func FindIfPathExists(graph map[int][]int, start, end int) bool {
	visited := make(map[int]bool)
	stack := []int{start}
	visited[start] = true

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == end {
			return true
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return false
}

// Example usage:
// graph := CreateGraphAdjacencyList([][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}})
// exists := FindIfPathExists(graph, 1, 4)
// fmt.Println("Path exists from 1 to 4:", exists) // Output: true
// exists = FindIfPathExists(graph, 1, 5)
// fmt.Println("Path exists from 1 to 5:", exists) // Output: false
// Note: This function checks if there is a path from 'start' to 'end' in the graph.
