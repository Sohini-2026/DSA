package graph

func TopologicalSortUsingBFS(graph map[int][]int) []int {
	inDegree := FindIndegree(graph)

	queue := []int{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	sortedOrder := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// we can check len of sortedOrder here to see if it matches the number of nodes in the graph
	if len(sortedOrder) != len(graph) {
		return []int{} // Cycle detected or not all nodes are reachable
	}
	return sortedOrder
}

func FindIndegree(graph map[int][]int) map[int]int {
	inDegree := make(map[int]int)
	for node := range graph {
		inDegree[node] = 0
	}

	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	return inDegree
}
