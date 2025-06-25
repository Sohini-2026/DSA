package binarySearch

func FindMinDiffElementInSortedArray(arr []int, target int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] < target {
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			return arr[mid] // Exact match found
		}
	}

	absdiffStart := abs(arr[start] - target)
	absDiffEnd := abs(arr[end] - target)

	if absdiffStart < absDiffEnd {
		return arr[start]
	} else if absDiffEnd < absdiffStart {
		return arr[end]
	}

	if absdiffStart == absDiffEnd {
		return arr[end] // If both are equal, return the smaller element
	}
	return -1 // This case should not happen if the input is valid
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
