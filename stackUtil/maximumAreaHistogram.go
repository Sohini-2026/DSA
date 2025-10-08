package stackUtil

/*
Find the largest rectangular area possible in a given histogram where the largest rectangle can be made of a number of contiguous bars. For simplicity, assume that all bars have same width and the width is 1 unit.
*/

/*
NSL = stack.Top()
NSR = i-1
Width = NSR - NSL
*/

import (
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func MaximumAreaHistogram(heights []int) int {
	stack := &baseDS.Stack{}
	maxArea := 0
	n := len(heights)

	for i := 0; i <= n; i++ {
		var currHeight int
		if i == n {
			currHeight = 0
		} else {
			currHeight = heights[i]
		}

		for !stack.IsEmpty() && currHeight < heights[stack.Top()] {
			topIndex := stack.Pop()
			height := heights[topIndex]
			var width int
			if stack.IsEmpty() {
				width = i
			} else {
				width = i - stack.Top() - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
			fmt.Printf("i: %d,Popped index: %d, Height: %d, Width: %d, Area: %d, MaxArea: %d\n", i, topIndex, height, width, area, maxArea)
		}
		stack.Push(i)
		fmt.Printf("Pushed index: %d, Current Stack: %+v\n", i, stack.Items)
	}

	return maxArea
}
