package graph

import "fmt"

/*
You are a hiker preparing for an upcoming hike. You are given heights, a 2D array of size rows x columns, where heights[row][col] represents the height of cell (row, col). You are situated in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). You can move up, down, left, or right, and you wish to find a route that requires the minimum effort.

A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.

Return the minimum effort required to travel from the top-left cell to the bottom-right cell.

Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
Output: 2
Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.
*/

func MinimumEffortPath(heights [][]int) int {
	rows := len(heights)
	cols := len(heights[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up

	// Create graph as adjacency list
	adjList := make(map[int][][]int)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			node := r*cols + c
			for _, dir := range directions {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < rows && newC >= 0 && newC < cols {
					neighbor := newR*cols + newC
					weight := abs(heights[r][c] - heights[newR][newC])
					adjList[node] = append(adjList[node], []int{neighbor, weight})
				}
			}
		}
	}

	start := 0
	target := rows*cols - 1

	fmt.Printf("Adjacency List: %+v\n", adjList)
	path, minEffort := Dijkstra(adjList, start, target)
	fmt.Println("Path:", path)
	return minEffort
}

func Dijkstra(adjList map[int][][]int, start int, target int) ([]int, int) {
	distances := make(map[int]int)
	visited := make(map[int]bool)
	prev := make(map[int]int)

	for node := range adjList {
		distances[node] = 1<<31 - 1 // Initialize distances to infinity
	}
	distances[start] = 0

	for len(visited) < len(adjList) {
		minNode := -1
		minDist := 1<<31 - 1
		for node, dist := range distances {
			if !visited[node] && dist < minDist {
				minDist = dist
				minNode = node
			}
		}

		if minNode == -1 {
			break
		}

		visited[minNode] = true

		for _, neighbor := range adjList[minNode] {
			nextNode, weight := neighbor[0], neighbor[1]
			if !visited[nextNode] {
				newDist := max(distances[minNode], weight)
				if newDist < distances[nextNode] {
					distances[nextNode] = newDist
					prev[nextNode] = minNode
				}
			}
		}
	}

	// Reconstruct path from start to target
	path := []int{}
	for at := target; at != start; at = prev[at] {
		path = append([]int{at}, path...)
		if _, ok := prev[at]; !ok {
			return nil, -1 // No path found
		}
	}
	path = append([]int{start}, path...)

	return path, distances[target]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
