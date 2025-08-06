package graph

import "sort"

func FindSafeNodes(graph map[int][]int) []int {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)
	safeNodes := []int{}

	for node := range graph {
		if !visited[node] {
			dfsDGCycle(graph, node, visited, recStack)

		}
	}

	// Filter out nodes that are part of cycles
	for nodeEntry, isInCycle := range recStack {
		if !isInCycle {
			safeNodes = append(safeNodes, nodeEntry)
		}
	}

	sort.Ints(safeNodes)

	return safeNodes
}
