package binarySearch

func FindElementInRowColumnSortedMatrix(matrix [][]int, target int) (int, int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1, -1 // Matrix is empty
	}

	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1

	for i >= 0 && i < n && j >= 0 && j < m {
		if matrix[i][j] == target {
			return i, j // Target found at (i, j)
		} else if matrix[i][j] > target {
			j-- // Move left
		} else {
			i++ // Move down
		}

	}

	return -1, -1 // Target not found
}
