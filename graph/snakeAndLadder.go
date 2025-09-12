package graph

// You are given an n x n integer matrix board where the cells are labeled from 1 to n2 in a Boustrophedon style starting from the bottom left of the board (i.e. board[n - 1][0]) and alternating direction each row.

// You start on square 1 of the board. In each move, starting from square curr, do the following:

// Choose a destination square next with a label in the range [curr + 1, min(curr + 6, n2)].
// This choice simulates the result of a standard 6-sided die roll: i.e., there are always at most 6 destinations, regardless of the size of the board.
// If next has a snake or ladder, you must move to the destination of that snake or ladder. Otherwise, you move to next.
// The game ends when you reach the square n2.
// A board square on row r and column c has a snake or ladder if board[r][c] != -1. The destination of that snake or ladder is board[r][c]. Squares 1 and n2 are not the starting points of any snake or ladder.

// Note that you only take a snake or ladder at most once per dice roll. If the destination to a snake or ladder is the start of another snake or ladder, you do not follow the subsequent snake or ladder.

// For example, suppose the board is [[-1,4],[-1,3]], and on the first move, your destination square is 2. You follow the ladder to square 3, but do not follow the subsequent ladder to 4.

func SnakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n
	visited := make([]bool, target+1)
	queue := [][2]int{{1, 0}} // {square, moves}
	visited[1] = true

	for len(queue) > 0 {
		curr, moves := queue[0][0], queue[0][1]
		queue = queue[1:]

		if curr == target {
			return moves
		}

		for die := 1; die <= 6; die++ {
			next := curr + die
			if next > target {
				break
			}
			r, c := getCoordinates(next, n)
			if board[r][c] != -1 {
				next = board[r][c]
			}
			if !visited[next] {
				visited[next] = true
				queue = append(queue, [2]int{next, moves + 1})
			}
		}
	}

	return -1 // If we cannot reach the target square
}

func getCoordinates(square, n int) (int, int) {
	square-- // Convert to 0-based index
	row := n - 1 - square/n
	col := square % n
	if (n-row)%2 == 0 {
		col = n - 1 - col
	}

	return row, col
}
