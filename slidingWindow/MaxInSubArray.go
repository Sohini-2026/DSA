package slidingwindow

import "fmt"

func MaxInSubArray(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return []int{} // Invalid input
	}

	result := make([]int, 0)
	queue := make([]int, 0) // To store indices of maximum elements

	for i := 0; i < len(arr); i++ {
		// Remove elements not in the current window
		if len(queue) > 0 && queue[0] < i-k+1 {
			queue = queue[1:]
		}

		// Remove elements smaller than the current element from the back of the queue
		for len(queue) > 0 && arr[queue[len(queue)-1]] < arr[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		// If we have filled the first window, add the result
		if i >= k-1 {
			result = append(result, arr[queue[0]])
		}
	}

	return result
}

func MaxInSubArray1(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return []int{} // Invalid input
	}

	result := make([]int, 0)
	i, j := 0, 0
	queue := make([]int, 0) // To store indices of maximum elements

	for j < len(arr) {
		// Remove indices out of the window from the front
		if len(queue) > 0 && queue[0] < i {
			queue = queue[1:]
		}
		// Remove elements smaller than the current element from the back
		for len(queue) > 0 && arr[queue[len(queue)-1]] < arr[j] {
			queue = queue[:len(queue)-1]
		}

		// Append current index
		queue = append(queue, j)

		fmt.Printf("Current queue: %+v\n", queue)

		// If window size is k, append the max to result
		if j-i+1 == k {
			result = append(result, arr[queue[0]])
			i++
		}
		j++
	}

	return result
}
