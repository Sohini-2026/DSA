package stackUtil

import (
	"github.com/Sohini-2026/DSA/baseDS"
)

/*Given an array, print the Next Greater Element (NGE) for every element. The Next greater Element for an element x is the first greater element on the right side of x in array. Elements for which no greater element exist, consider next greater element as -1.*/
func NextSmallerToLeft(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	stack := baseDS.Stack{}

	for i := 0; i < n; i++ {
		// Pop elements from the stack until we find a greater element or the stack becomes empty
		for stack.Size() > 0 && stack.Top() >= arr[i] {
			stack.Pop()
		}

		// If the stack is empty, it means there is no greater element to the right
		if stack.Size() == 0 {
			result[i] = -1
		} else {
			result[i] = stack.Top()
		}

		// Push the current element onto the stack
		stack.Push(arr[i])
	}

	return result
}
