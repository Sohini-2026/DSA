package backtracking

/*
Consider a rat placed at (0, 0) in a square matrix of order N * N. It has to reach the destination at (N - 1, N - 1). Find all possible paths that the rat can take to reach from source to destination. The directions in which the rat can move are 'U'(up), 'D'(down), 'L' (left), 'R' (right). Value 0 at a cell in the matrix represents that it is blocked and rat cannot move to it while value 1 at a cell in the matrix represents that rat can be travel through it.
Note: In a path, no cell can be visited more than one time. If the source cell is 0, the rat cannot move to any other cell.

Example 1:

Input:
N = 4
m[][] = {{1, 0, 0, 0},
         {1, 1, 0, 1},
         {1, 1, 0, 0},
         {0, 1, 1, 1}}
Output:
DDRDRR DRDDRR
Explanation:
The rat can reach the destination at
(3, 3) from (0, 0) by two paths - DRDDRR
and DDRDRR, when printed in sorted order
we get DDRDRR DRDDRR
*/

func FindPathInMaze(maze [][]int, n int) []string {
	var res []string
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	if maze[0][0] == 1 {
		backtrackMaze(maze, n, 0, 0, "", &res, visited)
	}
	return res
}

var directions = [][]int{
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
	{-1, 0}, // Up
}

var dirChars = []byte{'D', 'L', 'R', 'U'}

func backtrackMaze(maze [][]int, n, x, y int, path string, res *[]string, visited [][]bool) {
	if x == n-1 && y == n-1 {
		*res = append(*res, path)
		return
	}

	for i, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]

		if newX >= 0 && newX < n && newY >= 0 && newY < n &&
			maze[newX][newY] == 1 && !visited[newX][newY] {
			visited[newX][newY] = true
			backtrackMaze(maze, n, newX, newY, path+string(dirChars[i]), res, visited)
			visited[newX][newY] = false // backtrack
		}
	}
}

/*
The backtrack step visited[newX][newY] = false ensures that after exploring all paths from the current cell, we mark it as unvisited again.
This allows other potential paths to use this cell in their exploration.
*/
