package binarySearch

func FindElementInBitonicArray(arr []int, target int) int {
	peak := FindPeakElement(arr)
	if peak == -1 {
		return -1 // No peak found, array is not bitonic
	}

	// Search in the increasing part of the array
	index := FindNumberInSortedArray(arr, target, 0, peak, true)
	if index != -1 {
		return index
	}

	// Search in the decreasing part of the array
	return FindNumberInSortedArray(arr, target, peak+1, len(arr)-1, false)
}

func FindNumberInSortedArray(arr []int, target int, start int, end int, isAscending bool) int {
	left, right := start, end

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Target found
		} else if isAscending {
			if arr[mid] < target {
				left = mid + 1 // Search in the right half
			} else {
				right = mid - 1 // Search in the left half
			}
		} else {
			if arr[mid] > target {
				left = mid + 1 // Search in the right half
			} else {
				right = mid - 1 // Search in the left half
			}
		}
	}

	return -1 // Target not found
}
