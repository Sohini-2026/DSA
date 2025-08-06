package graph

import "fmt"

func IsCycleInUndirectedGraph(graph map[int][]int) bool {
	visited := make(map[int]bool)

	for node := range graph {
		if !visited[node] {
			if dfsDetectCycle(graph, node, visited, -1) {
				return true // Cycle found
			}
		}
	}

	return false // No cycle found
}

func dfsDetectCycle(graph map[int][]int, node int, visited map[int]bool, parent int) bool {
	visited[node] = true

	for _, neighbor := range graph[node] {
		fmt.Print("Visiting neighbor:", neighbor, "of node:", node, "\n", "belonging to parent:", parent, "\n")
		if !visited[neighbor] {
			if dfsDetectCycle(graph, neighbor, visited, node) {
				return true
			}
		} else if neighbor != parent {
			return true // Found a back edge
		}
	}

	return false
}

//AdjList : 1{2,5},2,{1,3},3{2,4},4{3,5},5{1,4}
// Visiting neighbor:2of node:1
// belonging to parent:-1
// Visiting neighbor:1of node:2
// belonging to parent:1
// Visiting neighbor:3of node:2
// belonging to parent:1
// Visiting neighbor:2of node:3
// belonging to parent:2
// Visiting neighbor:4of node:3
// belonging to parent:2
// Visiting neighbor:3of node:4
// belonging to parent:3
// Visiting neighbor:5of node:4
// belonging to parent:3
// Visiting neighbor:4of node:5
// belonging to parent:4
// Visiting neighbor:1of node:5
// belonging to parent:4

// As DFS we check in 1 direction first so 1->2->3->4->5->1
