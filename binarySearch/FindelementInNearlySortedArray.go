package binarySearch

// FindElementInNearlySortedArray finds an element in a nearly sorted array
// the element can be in i,i+1 or i-1 index of the mid element
func FindElementInNearlySortedArray(arr []int, target int) int {
	if len(arr) == 0 {
		return -1 // Array is empty
	}

	start, end := 0, len(arr)-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid // Target found at mid index
		}

		// Check if the target is at mid-1 or mid+1
		if (mid-1) > start && arr[mid-1] == target {
			return mid - 1 // Target found at mid-1 index
		}
		if (mid+1) < end && arr[mid+1] == target {
			return mid + 1 // Target found at mid+1 index
		}

		// Adjust search range based on the value at mid
		if arr[mid] < target {
			start = mid + 2 // Skip the next element since it's already checked
		} else {
			end = mid - 2 // Skip the previous element since it's already checked
		}
	}

	return -1 // Target not found in the array
}
