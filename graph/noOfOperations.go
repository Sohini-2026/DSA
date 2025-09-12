package graph

// There are n computers numbered from 0 to n - 1 connected by ethernet cables connections forming a network where connections[i] = [ai, bi] represents a connection between computers ai and bi. Any computer can reach any other computer directly or indirectly through the network.

// You are given an initial computer network connections. You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected.

// Return the minimum number of times you need to do this in order to make all the computers connected. If it is not possible, return -1.

// Example 1:

// Input: n = 4, connections = [[0,1],[0,2],[1,2]]
// Output: 1
// Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.

func MakeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1 // Not enough cables to connect all computers
	}

	adjList := make([][]int, n)
	for _, conn := range connections {
		adjList[conn[0]] = append(adjList[conn[0]], conn[1])
		adjList[conn[1]] = append(adjList[conn[1]], conn[0])
	}

	visited := make([]bool, n)
	componentCount := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			componentCount++
			dfsConnections(i, visited, adjList)
		}
	}

	return componentCount - 1 // Minimum operations needed is number of components - 1
}

func dfsConnections(node int, visited []bool, adjList [][]int) {
	visited[node] = true
	for _, neighbor := range adjList[node] {
		if !visited[neighbor] {
			dfsConnections(neighbor, visited, adjList)
		}
	}
}
