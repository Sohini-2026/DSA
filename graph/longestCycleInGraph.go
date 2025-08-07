package graph

import "fmt"

func FindLongestcycle(graph map[int][]int) int {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)
	longestCycle := 0

	for node := range graph {
		if !visited[node] {
			cycleLength := dfsFindLongestCycle(graph, node, visited, recStack, make(map[int]int), 0)
			if cycleLength > longestCycle {
				longestCycle = cycleLength
			}
		}
	}

	return longestCycle
}
func dfsFindLongestCycle(graph map[int][]int, node int, visited map[int]bool, recStack map[int]bool, pathIndex map[int]int, length int) int {
	visited[node] = true
	recStack[node] = true
	pathIndex[node] = length
	maxCycleLength := 0

	fmt.Printf("Visiting node: %d, current path length: %d\n,pathIndex::%+v", node, length, pathIndex)

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			cycleLength := dfsFindLongestCycle(graph, neighbor, visited, recStack, pathIndex, length+1)
			if cycleLength > maxCycleLength {
				maxCycleLength = cycleLength
			}
		} else if recStack[neighbor] {
			cycleLength := length - pathIndex[neighbor] + 1
			fmt.Printf("Cycle detected from node %d to %d with length %d\n", neighbor, node, cycleLength)
			if cycleLength > maxCycleLength {
				maxCycleLength = cycleLength
			}
		}
	}

	recStack[node] = false
	delete(pathIndex, node)
	return maxCycleLength
}
