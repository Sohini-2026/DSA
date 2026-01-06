package week1

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	originalColor := image[sr][sc]

	if originalColor == newColor {
		return image // No need to fill if the color is the same
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

		image[r][c] = newColor
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			queue = append(queue, [2]int{nr, nc})
		}
	}

	return image
}
