package backtracking

/*
Given a number K and string str of digits denoting a positive integer, build the largest number possible by performing swap operations on the digits of str at most K times.


Example 1:

Input:
K = 4
str = "1234567"
Output:
7654321
Explanation:
Three swaps can make the
input 1234567 to 7654321, swapping 1
with 7, 2 with 6 and finally 3 with 5
*/

var maxNum string

func FindLargestNumberAfterKSwaps(str string, k int) string {
	maxNum = str
	findMaximumNumUsingBT(str, k, 0)
	return maxNum
}

func findMaximumNumUsingBT(str string, k, start int) {
	if k == 0 || start == len(str)-1 {
		return
	}
	maxChar := str[start]
	for i := start + 1; i < len(str); i++ {
		if str[i] > maxChar {
			maxChar = str[i]
		}
	}

	//fmt.Println("Maxchar::::::", maxChar)

	for i := start + 1; i < len(str); i++ {
		if str[start] < str[i] && str[i] == maxChar {
			//fmt.Println("Swapping ", str[start], " and ", str[i])
			str = swap(str, start, i)
			if str > maxNum {
				maxNum = str
			}
			findMaximumNumUsingBT(str, k-1, start+1)
			str = swap(str, start, i) // backtrack
		}
	}

	findMaximumNumUsingBT(str, k, start+1)
}

func findMaximumNum(str string, k int) {
	if k == 0 {
		return
	}
	n := len(str)
	for i := 0; i < n-1; i++ {
		// Find the max digit to the right of i
		maxDigit := str[i]
		for j := i + 1; j < n; j++ {
			if str[j] > maxDigit {
				maxDigit = str[j]
			}
		}
		// Only swap if a greater digit exists
		if maxDigit != str[i] {
			for j := i + 1; j < n; j++ {
				if str[j] == maxDigit {
					swapped := swap(str, i, j)
					if swapped > maxNum {
						maxNum = swapped
					}
					findMaximumNum(swapped, k-1)
					// No need to backtrack string, as strings are immutable in Go
				}
			}
			// Only need to process the first position where a swap is possible
			break
		}
	}
}

func swap(str string, i, j int) string {
	runes := []rune(str)
	runes[i], runes[j] = runes[j], runes[i]
	return string(runes)
}

// Time Complexity: O((n^2)^k) where n is the length of the string and k is the number of swaps allowed.
// Auxiliary Space: O(n) where n is the length of the string.
