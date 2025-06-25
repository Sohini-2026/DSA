package binarySearch

import "fmt"

func FindPositionOfElementInInfiniteSortedArray(arr []int, target int) int {
	start, end := 0, 1

	// Find the range where the target might be present
	for arr[end] < target {
		start = end
		end *= 2

		fmt.Println("Expanding range:", start, "to", end)
	}

	// Perform binary search in the identified range
	for start <= end {
		mid := start + (end-start)/2

		if mid >= len(arr) {
			return -1 // Out of bounds
		}

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1 // Element not found
}
