package slidingwindow

// FindFirstNegativeNumberInAllWindowOfSizeK finds the first negative number in each sliding window of size k.
func FindFirstNegativeNumberInAllWindowOfSizeK(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return []int{} // Invalid input
	}

	result := make([]int, 0)
	queue := make([]int, 0) // To store indices of negative numbers

	for i := 0; i < len(arr); i++ {
		// Remove elements not in the current window
		if len(queue) > 0 && queue[0] < i-k+1 {
			queue = queue[1:]
		}

		// Add current element if it is negative
		if arr[i] < 0 {
			queue = append(queue, i)
		}

		// If we have filled the first window, add the result
		if i >= k-1 {
			if len(queue) > 0 {
				result = append(result, arr[queue[0]])
			} else {
				result = append(result, 0) // No negative number in this window
			}
		}
	}

	return result
}

func FindFirstNegativeNumberInAllWindowOfSizeK1(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return []int{} // Invalid input
	}

	result := make([]int, 0)
	queue := make([]int, 0) // To store indices of negative numbers

	i, j := 0, 0

	for j < len(arr) {
		// Add current element if it is negative
		if arr[j] < 0 {
			queue = append(queue, j)
		}

		// If we have filled the first window, add the result
		if j-i+1 == k {
			if len(queue) > 0 && queue[0] >= i {
				result = append(result, arr[queue[0]])
			} else {
				result = append(result, 0) // No negative number in this window
			}
			i++

			if len(queue) > 0 && queue[0] < i {
				queue = queue[1:]
			}
		}
		j++
	}

	return result
}

// Brute force approach to find first negative number in all windows of size k
func FindFirstNegativeNumberInAllWindowOfSizeKBruteForce(arr []int, k int) []int {
	result := make([]int, 0)
	n := len(arr)
	for i := 0; i <= n-k; i++ {
		found := false
		for j := i; j < i+k; j++ {
			if arr[j] < 0 {
				result = append(result, arr[j])
				found = true
				break
			}
		}
		if !found {
			result = append(result, 0)
		}
	}
	return result
}
