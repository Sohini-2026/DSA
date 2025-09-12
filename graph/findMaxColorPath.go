package graph

import "fmt"

// There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.

// You are given a string colors where colors[i] is a lowercase English letter representing the color of the ith node in this graph (0-indexed). You are also given a 2D array edges where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.

// A valid path in the graph is a sequence of nodes x1 to x2 to x3 to ... to xk such that there is a directed edge from xi to xi+1 for every 1 less than or = i less than k. The color value of the path is the number of nodes that are colored the most frequently occurring color along that path.

// Return the largest color value of any valid path in the given graph, or -1 if the graph contains a cycle.

// Example 1:
// Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
// Output: 3
// Explanation: The path 0 to 2 to 3 to 4 contains 3 nodes that are colored "a" (red in the above image).

func FindLargestColorValue(colors string, edges [][]int) int {
	graph := make(map[int][]int)
	inDegree := make([]int, len(colors))
	colorCount := make([]map[byte]int, len(colors))

	// Build the graph and in-degree count
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		inDegree[b]++
	}

	// Initialize color count maps
	for i := range colorCount {
		colorCount[i] = make(map[byte]int)
		colorCount[i][colors[i]] = 1
	}

	fmt.Println("Color Count Initialization:", colorCount)

	queue := []int{}

	// Start with nodes that have no incoming edges
	for i, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, i)
		}
	}

	maxColorValue := -1
	visitedCount := 0 // Track number of processed nodes

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visitedCount++
		maxFreq := 0
		for _, freq := range colorCount[node] {
			if freq > maxFreq {
				maxFreq = freq
			}
		}

		maxColorValue = max(maxColorValue, maxFreq)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--

			for color, count := range colorCount[node] {
				colorCount[neighbor][color] += count
			}

			fmt.Println("Processing node:", node, "Neighbor:", neighbor, "Color Count:", colorCount[node])

			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if visitedCount < len(colors) {
		return -1 // Cycle detected
	}

	return maxColorValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Note: This function assumes that the input graph does not contain cycles. If cycles are present, additional cycle detection logic would be needed.
// The function returns -1 if a cycle is detected, but this implementation does not include cycle detection logic.
// To handle cycles, you would typically use a visited set to track nodes during DFS or BFS and check for back edges.
// If a cycle is detected, you would return -1.
// This implementation focuses on finding the maximum color value along valid paths in a directed acyclic graph (DAG).
// If you need to handle cycles, you can modify the function to include cycle detection logic.
// For example, you can use a visited set to track nodes during traversal and check for back edges.
// If a back edge is found, you can return -1 to indicate that a cycle exists in the graph.
// This implementation assumes that the input graph is a directed acyclic graph (DAG) and does not include cycle detection logic.
