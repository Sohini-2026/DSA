package graph

import "fmt"

func CreateDirectedGraphAdjacencyList(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
	}
	return graph
}

func IsCycleInDirectedGraphByDFS(graph map[int][]int) bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for node := range graph {
		fmt.Printf("Checking node: %d\n", node)
		if !visited[node] {
			if dfsDGCycle(graph, node, visited, recStack) {
				return true // Cycle found
			}
		}
	}

	return false // No cycle found
}
func dfsDGCycle(graph map[int][]int, node int, visited map[int]bool, recStack map[int]bool) bool {
	visited[node] = true
	recStack[node] = true

	for _, neighbor := range graph[node] {
		fmt.Printf("Visiting neighbor: %d of node: %d\n and recStack :: %+v and visited :: %+v", neighbor, node, recStack, visited)
		if !visited[neighbor] {
			if dfsDGCycle(graph, neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			return true // Found a back edge indicating a cycle
		}

	}

	recStack[node] = false // Remove the node from recursion stack

	return false
}
