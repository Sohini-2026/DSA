package graph

func DepthFirstTraversal(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	stack := []int{start}
	visited[start] = true
	result := []int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return result
}

// Example usage:
// graph := CreateGraphAdjacencyList([][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}})
// dfsResult := DFS(graph, 1)
// fmt.Println(dfsResult) // Output: [1 3 4 2] or similar order depending on the graph structure
// Note: The order of traversal may vary based on the graph structure and the order of neighbors in the adjacency list.
// This DFS implementation assumes an undirected graph. For directed graphs, you can remove the line that appends the reverse edge.
// This DFS implementation uses an iterative approach with a stack to avoid recursion depth issues in large graphs.
