package backtracking

/*
The n-queens puzzle is the problem of placing n queens on a (n×n) chessboard such that no two queens can attack each other.
Given an integer n, find all distinct solutions to the n-queens puzzle. Each solution contains distinct board configurations of the n-queens placement, where the solutions are a permutation of [1,2,3..n] in increasing order, here the number in the ith place denotes that the ith-column queen is placed in the row with that number.

Examples :

Input: 1
Output: [1]
Explaination: Only one queen can be placed
in the single cell available.
Input: 4
Output: [2 4 1 3 ] [3 1 4 2 ]
Explaination: These are the 2 possible solutions.
*/

func NQueens(n int) [][]int {
	var res [][]int
	var path []int
	cols := make(map[int]bool)
	diag1 := make(map[int]bool) // row - col
	diag2 := make(map[int]bool) // row + col
	backtrackNQueens(n, 0, path, cols, diag1, diag2, &res)
	return res
}

func backtrackNQueens(n, row int, path []int, cols, diag1, diag2 map[int]bool, res *[][]int) {
	if row == n {
		solution := make([]int, n)
		copy(solution, path)
		*res = append(*res, solution)
		return
	}

	for col := 0; col < n; col++ {
		if cols[col] || diag1[row-col] || diag2[row+col] {
			continue
		}

		path = append(path, col+1) // Store 1-based index
		cols[col] = true
		diag1[row-col] = true
		diag2[row+col] = true

		backtrackNQueens(n, row+1, path, cols, diag1, diag2, res)

		// Backtrack
		path = path[:len(path)-1]
		delete(cols, col)
		delete(diag1, row-col)
		delete(diag2, row+col)
	}
}
