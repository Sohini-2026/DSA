package binarySearch

func FindNextAlphabeticalElement(arr []string, target string) string {
	start, end := 0, len(arr)-1
	res := "#"
	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] <= target {
			start = mid + 1
		} else {
			res = arr[mid]
			end = mid - 1
		}
	}

	return res
}
