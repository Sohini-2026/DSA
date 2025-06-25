package binarySearch

// this is appilied in finding max element in bitonic array
func FindPeakElement(arr []int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := start + (end-start)/2

		if mid > 0 && mid < len(arr)-1 {
			if (arr[mid] > arr[mid+1]) && (arr[mid] > arr[mid-1]) {
				return mid // Found the peak element
			} else if arr[mid] < arr[mid+1] {
				start = mid + 1 // Move to the right side
			} else {
				end = mid - 1 // Move to the left side
			}
		} else if mid == 0 {
			if arr[mid] > arr[mid+1] {
				return mid // Found the peak element
			} else {
				return mid + 1 // Move to the right side
			}
		} else if mid == len(arr)-1 {
			if arr[mid] > arr[mid-1] {
				return mid // Found the peak element
			} else {
				return mid - 1 // Move to the left side
			}
		}
	}

	return -1 // or end, both will point to the peak element
}
