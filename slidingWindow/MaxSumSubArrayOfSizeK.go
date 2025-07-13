package slidingwindow

import "math"

// MaxSumSubArrayOfSizeK finds the maximum sum of a subarray of size k in an array.
func MaxSumSubArrayOfSizeK(arr []int, k int) int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return 0 // Invalid input
	}

	maxSum := math.MinInt64
	currentSum := 0

	// Calculate the sum of the first 'k' elements
	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}
	maxSum = currentSum

	// Slide the window over the rest of the array
	for i := k; i < len(arr); i++ {
		currentSum += arr[i] - arr[i-k] // Add next element and remove the first element of the previous window
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func MaxSumSubArrayOfSizeK1(arr []int, k int) int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return 0 // Invalid input
	}

	maxSum := math.MinInt64
	currentSum := 0

	i, j := 0, 0

	for j < len(arr) {
		currentSum += arr[j] // Add the next element to the current sum
		if j-i+1 == k {      // When we have a full window of size k
			maxSum = max(maxSum, currentSum) // Ensure maxSum is updated
			currentSum -= arr[i]             // Remove the first element of the window
			i++                              // Slide the window forward
		}
		j++ // Move to the next element
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
