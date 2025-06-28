package binarySearch

func AllocateMinNoOfPages(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return -1
	}

	start := max(arr) // The maximum number of pages in a single book
	end := sum(arr)   // The sum of all pages in the books
	res := -1
	for start <= end {
		mid := start + (end-start)/2

		if isValid(arr, n, k, mid) {
			res = mid // Found a valid allocation, try for a smaller maximum
			end = mid - 1
		} else {
			start = mid + 1 // Increase the minimum limit
		}
	}
	return res
}
func max(arr []int) int {
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func isValid(arr []int, n, k, mid int) bool {
	sum := 0
	student := 1

	for i := 0; i < n; i++ {
		if arr[i] > mid {
			return false // If a single book has more pages than mid, it's not possible
		}

		if sum+arr[i] > mid {
			student++        // Allocate to the next student
			sum = arr[i]     // Start new allocation with current book
			if student > k { // If we exceed the number of students, return false
				return false
			}
		} else {
			sum += arr[i] // Add to current student's allocation
		}
	}

	return true // Valid allocation found
}
