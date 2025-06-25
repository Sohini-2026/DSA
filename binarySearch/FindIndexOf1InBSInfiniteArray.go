package binarySearch

func FindIndexOf1InBSInfiniteArray(arr []int) int {
	start, end := 0, 1
	res := 0
	// Find the range where the first 1 might be present
	for arr[end] == 0 {
		start = end
		end *= 2
	}

	// Perform binary search in the identified range
	for start <= end {
		mid := start + (end-start)/2

		if mid >= len(arr) {
			return -1 // Out of bounds
		}

		if arr[mid] >= 1 {
			res = mid
			// Continue searching to the left for the first occurrence
			end = mid - 1
			// Found the first occurrence of 1
		} else if arr[mid] < 1 {
			start = mid + 1
		}
	}

	return res // Element not found
}
