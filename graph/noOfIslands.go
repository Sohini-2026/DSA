package graph

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

// Example 1:

// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1

func NoOfIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	islandCount := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 && !visited[r][c] {
				dfsIslandCount(r, c, rows, cols, grid, visited)
				islandCount++
			}
		}
	}

	return islandCount
}

func dfsIslandCount(r, c, rows, cols int, grid [][]int, visited [][]bool) {
	if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 || visited[r][c] {
		return
	}
	visited[r][c] = true
	dfsIslandCount(r+1, c, rows, cols, grid, visited) // Down
	dfsIslandCount(r-1, c, rows, cols, grid, visited) // Up
	dfsIslandCount(r, c+1, rows, cols, grid, visited) // Right
	dfsIslandCount(r, c-1, rows, cols, grid, visited) // Left
}
