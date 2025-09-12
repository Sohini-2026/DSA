package graph

// You are given an image represented by an m x n grid of integers image, where image[i][j] represents the pixel value of the image. You are also given three integers sr, sc, and color. Your task is to perform a flood fill on the image starting from the pixel image[sr][sc].

// To perform a flood fill:

// Begin with the starting pixel and change its color to color.
// Perform the same process for each pixel that is directly adjacent (pixels that share a side with the original pixel, either horizontally or vertically) and shares the same color as the starting pixel.
// Keep repeating this process by checking neighboring pixels of the updated pixels and modifying their color if it matches the original color of the starting pixel.
// The process stops when there are no more adjacent pixels of the original color to update.
// Return the modified image after performing the flood fill.

// Example 1:

// Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2

// Output: [[2,2,2],[2,2,0],[2,0,1]]

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image // No need to fill if the color is the same
	}

	rows, cols := len(image), len(image[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up

	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || image[r][c] != originalColor {
			return
		}
		image[r][c] = color // Change the color
		for _, dir := range directions {
			dfs(r+dir[0], c+dir[1]) // Explore neighbors
		}
	}

	dfs(sr, sc)
	return image
}

func FloodFillBFS(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}

	rows, cols := len(image), len(image[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up
	queue := [][2]int{{sr, sc}}

	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		if r < 0 || r >= rows || c < 0 || c >= cols || image[r][c] != originalColor {
			continue
		}

		image[r][c] = color
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			queue = append(queue, [2]int{nr, nc})
		}
	}

	return image
}
