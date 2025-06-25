package binarySearch

// Logic:
// 7.8 -> here 7 is floor and 8 is ceil
// Floor->{1 2 4 8 9 10},5 if 5 is present return but if not return greatest element in array nearest to 5 => 4
// Ceil->{1 2 4 8 9 10},5 if 5 is present return but if not return lowest element in array nearest to 5 =>8
// FindFloorAndCeilOfElementInSortedArray finds the floor and ceil of a given element in a sorted array.
func FindFloorOfElementInSortedArray(arr []int, target int) int {
	start, end := 0, len(arr)
	res := 0
	for start <= end {
		mid := start + (end-start)/2

		if target == arr[mid] {
			return mid
		}

		if arr[mid] < target {
			res = mid
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		}
	}

	return res

}

func FindCeilOfElementInSortedArray(arr []int, target int) int {
	start, end := 0, len(arr)
	res := 0
	for start <= end {
		mid := start + (end-start)/2

		if target == arr[mid] {
			return mid
		}

		if arr[mid] < target {
			start = mid + 1
		} else if arr[mid] > target {
			res = mid
			end = mid - 1
		}
	}

	return res
}
