package graph

/*
There are n cities connected by some number of flights. You are given an array flights where flights[i] = [fromi, toi, pricei] indicates that there is a flight from city fromi to city toi with cost pricei.

You are also given three integers src, dst, and k, return the cheapest price from src to dst with at most k stops. If there is no such route, return -1.

Example 1:
Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
Output: 700
Explanation:
The graph is shown above.
The optimal path with at most 1 stop from city 0 to 3 is marked in red and has cost 100 + 600 = 700.
Note that the path through cities [0,1,2,3] is cheaper but is invalid because it uses 2 stops.
*/
func CheapestFlight(n int, flights [][]int, src int, dst int, k int) int {
	// Create adjacency list
	adjList := make(map[int][][]int)
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		adjList[from] = append(adjList[from], []int{to, price})
	}

	type NodeInfo struct {
		city      int
		cost      int
		stopsLeft int
	}

	queue := []NodeInfo{{src, 0, k + 1}} // k+1 because we can take k stops means k+1 edges

	minCost := make([]int, n)
	for i := range minCost {
		minCost[i] = 1<<31 - 1 // Initialize to infinity
	}
	minCost[src] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.city == dst {
			continue // We reached the destination, no need to explore further from here
		}

		if current.stopsLeft > 0 {
			for _, neighbor := range adjList[current.city] {
				nextCity, price := neighbor[0], neighbor[1]
				newCost := current.cost + price
				if newCost < minCost[nextCity] {
					minCost[nextCity] = newCost
					queue = append(queue, NodeInfo{nextCity, newCost, current.stopsLeft - 1})
				}
			}
		}
	}

	if minCost[dst] == 1<<31-1 {
		return -1 // No valid route found
	}
	return minCost[dst]
}
