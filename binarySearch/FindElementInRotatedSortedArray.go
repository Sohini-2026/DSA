package binarySearch

func FindElementInRotatedSortedArray(arr []int, target int) int {

	minIndex := CountNoOfTimesArrayIsRotated(arr)

	res1 := searchNumberInSortedArray(arr, target, 0, minIndex-1)
	if res1 > 0 {
		return res1 // Target found in the first half
	}
	res2 := searchNumberInSortedArray(arr, target, minIndex, len(arr)-1)
	if res2 > 0 {
		return res2 // Target found in the second half
	}

	return -1 // Target not found in either half
}

func searchNumberInSortedArray(arr []int, target int, start int, end int) int {
	left, right := start, end

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
