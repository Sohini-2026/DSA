package graph

var stack []int

func TopologicalSortUsingDFS(graph map[int][]int) []int {
	visited := make(map[int]bool)
	stack = []int{}

	for node := range graph {
		if !visited[node] {
			dfs(node, visited, graph)
		}
	}

	// Reverse the stack to get the topological order
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}
func dfs(node int, visited map[int]bool, graph map[int][]int) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, visited, graph)
		}
	}
	stack = append(stack, node)
}
