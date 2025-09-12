package graph

// Geek's village is represented by a 2-D matrix of characters of size n*m, where

// H - Represents a house
// W - Represents a well
// . - Represents an open ground
// N - Prohibited area(Not allowed to enter this area)

// Every house in the village needs to take water from a well, as the family members are so busy with their work, so every family wants to take the water from a well in minimum time, which is possible only if they have to cover as less distance as possible. Your task is to determine the minimum distance that a person in the house should travel to take out the water and carry it back to the house.

// A person is allowed to move only in four directions left, right, up, and down. That means if he/she is the cell (i,j), then the possible cells he/she can reach in one move are (i,j-1),(i,j+1),(i-1,j),(i+1,j), and the person is not allowed to move out of the grid.

// Note: For all the cells containing 'N', 'W' and '.' our answer should be 0, and for all the houses where there is no possibility of taking water our answer should be -1.

// Example 1:

// Input:
// n = 3
// m = 3
// c[][]: H H H
//        H W H
//        H H H
// Output:
// 4 2 4
// 2 0 2
// 4 2 4

func GeeksVillage(n int, m int, grid [][]byte) [][]int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up
	distance := make([][]int, n)
	visited := make([][]bool, n)
	queue := [][2]int{}

	for i := range distance {
		distance[i] = make([]int, m)
		visited[i] = make([]bool, m)
	}

	// Initialize the queue with all wells and mark them as visited
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 'W' {
				queue = append(queue, [2]int{r, c})
				visited[r][c] = true
			} else if grid[r][c] == 'N' || grid[r][c] == '.' {
				distance[r][c] = 0
				visited[r][c] = true
			} else {
				distance[r][c] = -1 // Initialize houses with -1
			}
		}
	}

	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < n && newC >= 0 && newC < m && !visited[newR][newC] && grid[newR][newC] == 'H' {
				visited[newR][newC] = true
				distance[newR][newC] = distance[r][c] + 2 // Distance to well and back
				queue = append(queue, [2]int{newR, newC})
			}
		}
	}

	return distance
}
