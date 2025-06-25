package binarySearch

func SearchNumberInSortedArray(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Target found
		} else if arr[mid] < target {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}

	return -1 // Target not found
}

func SearchNumberInReverseSortedArray(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Target found
		} else if arr[mid] > target {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}

	return -1 // Target not found
}
