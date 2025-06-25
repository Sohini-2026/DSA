package binarySearch

// Logic:
// Find index of min element.
// If not mid then search in unsorted array
func CountNoOfTimesArrayIsRotated(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	start, end := 0, len(arr)-1

	if arr[start] < arr[end] {
		return 0 // Array is not rotated
	}

	min := 0

	for start < end {
		mid := start + (end-start)/2
		prev := (mid + len(arr) - 1) % len(arr) // Previous index in circular array
		next := (mid + 1) % len(arr)            // Next index in circular array

		if arr[mid] < arr[prev] && arr[mid] < arr[next] {
			min = len(arr) - mid // Found the minimum element
			return min
		}

		if arr[mid] < arr[start] {
			end = mid // Search in the left half
		} else if arr[mid] > arr[end] {
			start = mid + 1 // Search in the right half
		}

	}

	return min // Return the index of the minimum element
}

//Q. how many times the array is rotated if its completely rotated?
//{7, 6, 5, 4, 3, 2, 1}
