package slidingwindow

func WindowsSubArrayOfSumK(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return []int{} // Invalid input
	}

	result := make([]int, 0)
	i, j := 0, 0
	currentSum := 0

	for j < len(arr) {
		currentSum += arr[j] // Add the next element to the current sum

		if currentSum == k { // When we have a full window of size k
			result = append(result, arr[i:j+1]...) // Append the current window
		}

		for currentSum >= k { // Shrink the window from the left if sum exceeds k
			currentSum -= arr[i]
			i++
		}
		j++ // Move to the next element
	}

	return result
}

func LargestWindowSubArrayOfSumK(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 {
		return []int{} // Invalid input
	}

	i, j := 0, 0
	currentSum := 0
	maxLen := 0
	startIdx := -1
	endIdx := -1

	for j < len(arr) {
		currentSum += arr[j]

		for currentSum > k && i <= j { // Shrink the window from the left if sum exceeds k
			currentSum -= arr[i]
			i++
		}

		if currentSum == k && (j-i+1) > maxLen { // Check if we found a larger window
			maxLen = j - i + 1
			startIdx = i
			endIdx = j
		}
		j++
	}

	if startIdx != -1 && endIdx != -1 {
		return arr[startIdx : endIdx+1]
	}

	// if we want to return length then return maxLen
	return []int{}
}
