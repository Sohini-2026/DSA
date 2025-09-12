package graph

// You are given an m x n grid where each cell can have one of three values:

// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

// Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

// Example 1:

// Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
// Output: 4
// Example 2:

// Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
// Output: -1
// Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

func RottingOranges(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	queue := [][2]int{}
	freshCount := 0

	// Initialize the queue with rotten oranges and count fresh oranges
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			} else if grid[r][c] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0 // No fresh oranges to rot
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up
	minutes := 0

	for len(queue) > 0 && freshCount > 0 {
		minutes++
		nextQueue := [][2]int{}

		for _, pos := range queue {
			r, c := pos[0], pos[1]

			for _, dir := range directions {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < rows && newC >= 0 && newC < cols && grid[newR][newC] == 1 {
					grid[newR][newC] = 2 // Rot the fresh orange
					freshCount--         // Decrease the count of fresh oranges
					nextQueue = append(nextQueue, [2]int{newR, newC})
				}
			}
		}

		queue = nextQueue // Move to the next level of BFS
	}

	// If all fresh oranges are rotten, return minutes; otherwise, return -1
	if freshCount == 0 {
		return minutes
	}
	return -1

}
