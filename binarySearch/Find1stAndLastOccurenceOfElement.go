package binarySearch

func FindFirstAndLastOccurrence(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1
	first, last := -1, -1

	// Find first occurrence
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			first = mid
			right = mid - 1 // Continue searching in the left half
		} else if arr[mid] < target {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}

	left, right = 0, len(arr)-1

	// Find last occurrence
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			last = mid
			left = mid + 1 // Continue searching in the right half
		} else if arr[mid] < target {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}

	return first, last
}

func FindCountOfOccurrences(arr []int, target int) int {
	first, last := FindFirstAndLastOccurrence(arr, target)
	if first == -1 || last == -1 {
		return 0 // Target not found
	}
	return last - first + 1 // Count of occurrences
}
