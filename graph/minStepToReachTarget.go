package graph

func MinStepByKnightToReachTarget(start, target [2]int) int {
	if start == target {
		return 0
	}

	directions := [][2]int{
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
	}

	queue := [][3]int{{start[0], start[1], 0}} // x, y, steps
	visited := make(map[[2]int]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		x, y, steps := current[0], current[1], current[2]

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			newPos := [2]int{newX, newY}

			if newPos == target {
				return steps + 1
			}

			if !visited[newPos] {
				visited[newPos] = true
				queue = append(queue, [3]int{newX, newY, steps + 1})
			}
		}
	}

	return -1 // If target is unreachable
}
