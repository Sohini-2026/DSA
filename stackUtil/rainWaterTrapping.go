package stackUtil

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
Input: arr[]   = {2, 0, 2}
Output: 2
Structure is like below
| |
|_|
We can trap 2 units of water in the middle gap.

Input: arr[]   = {3, 0, 0, 2, 0, 4}
Output: 10
Structure is like below
     |
|    |
|  | |
|__|_|
We can trap "3*2 units" of water between 3 an 2,
"1 unit" on top of bar 2 and "3 units" between 2
and 4.  See below diagram also.
*/

import (
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func RainWaterTrapping(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	stack := baseDS.Stack{}
	waterTrapped := 0

	for i := 0; i < len(heights); i++ {
		for !stack.IsEmpty() && heights[i] > heights[stack.Top()] {
			top := stack.Pop()
			if stack.IsEmpty() {
				break
			}
			distance := i - stack.Top() - 1
			boundedHeight := min(heights[i], heights[stack.Top()]) - heights[top]
			waterTrapped += distance * boundedHeight
			fmt.Printf("i: %d, top: %d, distance: %d, boundedHeight: %d, waterTrapped: %d\n", i, top, distance, boundedHeight, waterTrapped)
		}
		stack.Push(i)
		fmt.Printf("Stack after pushing index %d: %+v\n", i, stack)
	}

	return waterTrapped
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
