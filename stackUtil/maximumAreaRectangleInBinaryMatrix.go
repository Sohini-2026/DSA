package stackUtil

/*
Given a binary matrix, find the maximum size rectangle binary-sub-matrix with all 1’s.
Example:

Input :   0 1 1 0
          1 1 1 1
          1 1 1 1
          1 1 0 0

Output :  1 1 1 1
          1 1 1 1 .
*/

import (
	"fmt"
)

func MaximumAreaRectangleInBinaryMatrix(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	maxArea := 0
	rows := len(matrix)
	cols := len(matrix[0])
	histogram := make([]int, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				histogram[j] = 0
			} else {
				histogram[j] += 1
			}
		}
		fmt.Printf("Histogram for row %d: %+v\n", i, histogram)
		area := MaximumAreaHistogram(histogram)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
