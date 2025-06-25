package binarySearch

func SearchNumberWithUnknownSorting(arr []int, target int) int {

	if arr[0] < arr[1] {
		return SearchNumberInSortedArray(arr, target)
	} else {
		return SearchNumberInReverseSortedArray(arr, target)
	}
	// If the array is neither sorted nor reverse sorted, return -1
}
