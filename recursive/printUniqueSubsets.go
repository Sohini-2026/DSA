package recursive

import "sort"

func PrintUniqueSubsets(arr []int, index int, current []int, result *[][]int) {
	if index == len(arr) {
		// Sort to ensure unique subsets are printed in order
		sortedCurrent := make([]int, len(current))
		copy(sortedCurrent, current)
		sort.Ints(sortedCurrent)

		// Check if the subset is already present in the result
		for _, subset := range *result {
			if equal(subset, sortedCurrent) {
				return
			}
		}

		*result = append(*result, sortedCurrent)
		return
	}

	// Include the current element
	PrintUniqueSubsets(arr, index+1, append(current, arr[index]), result)

	// Exclude the current element
	PrintUniqueSubsets(arr, index+1, current, result)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func PrintUniqueSubsetsLexicographially(input, output string, m map[string]bool, result *[]string) {
	if len(input) == 0 {
		if !m[output] {
			m[output] = true
			*result = append(*result, output)
		}
		return
	}

	// Exclude the first character
	PrintUniqueSubsetsLexicographially(input[1:], output, m, result)
	// Include the first character
	PrintUniqueSubsetsLexicographially(input[1:], output+string(input[0]), m, result)
}
