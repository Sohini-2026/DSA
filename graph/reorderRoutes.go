package graph

import "fmt"

// There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree). Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.

// Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.

// This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

// Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.

// It's guaranteed that each city can reach city 0 after reorder.

func ReorderRoutes(connections [][]int, start int) int {
	// Create a graph representation of the roads
	adjList := make(map[int][]int)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		adjList[a] = append(adjList[a], b) // a -> b
		adjList[b] = append(adjList[b], a) // b -> a (for undirected graph)
	}

	fmt.Printf("Graph: %v\n", adjList)

	visited := make(map[int]bool)
	stack := []int{start} // Start from city 0
	visited[start] = true
	reorderCount := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, neighbor := range adjList[node] {
			fmt.Print("Checking neighbor: ", neighbor, " of node: ", node, "\n")
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)

				// Only increment if the original edge is from node to neighbor
				for _, conn := range connections {
					fmt.Println("Checking connection:", conn)
					if conn[0] == node && conn[1] == neighbor {
						fmt.Print("incrementing reorder count for connection: ", conn, "\n")
						reorderCount++
						break
					}
				}
			}
		}
	}

	return reorderCount
}

func ReorderRoutes1(connections [][]int, start int) int {
	adjList := make(map[int][]int)
	edgeDir := make(map[[2]int]bool)

	for _, conn := range connections {
		a, b := conn[0], conn[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
		edgeDir[[2]int{a, b}] = true // original direction
	}

	fmt.Printf("Graph: %v\n", adjList)
	fmt.Printf("Edge Directions: %v\n", edgeDir)

	visited := make(map[int]bool)
	stack := []int{start}
	visited[start] = true
	reorderCount := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, neighbor := range adjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
				// Increment only if the original edge is from node to neighbor
				if edgeDir[[2]int{node, neighbor}] {
					reorderCount++
				}
			}
		}
	}

	return reorderCount
}
