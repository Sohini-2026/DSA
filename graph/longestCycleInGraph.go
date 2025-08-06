package graph

func FindLongestcycle(graph map[int][]int) int {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)
	longestCycle := 0

	for node := range graph {
		if !visited[node] {
			cycleLength := dfsFindLongestCycle(graph, node, visited, recStack, 0)
			if cycleLength > longestCycle {
				longestCycle = cycleLength
			}
		}
	}

	return longestCycle
}
func dfsFindLongestCycle(graph map[int][]int, node int, visited map[int]bool, recStack map[int]bool, length int) int {
	visited[node] = true
	recStack[node] = true
	maxCycleLength := 0

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			cycleLength := dfsFindLongestCycle(graph, neighbor, visited, recStack, length+1)
			if cycleLength > maxCycleLength {
				maxCycleLength = cycleLength
			}
		} else if recStack[neighbor] {
			// Found a back edge indicating a cycle
			cycleLength := length + 1 // Include the current node in the cycle length
			if cycleLength > maxCycleLength {
				maxCycleLength = cycleLength
			}
		}
	}

	recStack[node] = false // Remove the node from recursion stack
	return maxCycleLength
}
