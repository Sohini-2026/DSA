package backtracking

/*
Given an integer N, print all the N digit numbers in increasing order, such that their digits are in strictly increasing order(from left to right).

Example 1:

Input:
N = 1
Output:
0 1 2 3 4 5 6 7 8 9
Explanation:
Single digit numbers are considered to be
strictly increasing order.
*/

import (
	"strconv"
)

func solve(n int, res *[]string, nums *int, last int) {
	if n == 0 {
		*res = append(*res, strconv.Itoa(*nums))
		return
	}
	for i := last + 1; i < 10; i++ {
		*nums = *nums*10 + i
		solve(n-1, res, nums, i)
		*nums /= 10 // backtrack
	}
}

func PrintNDigitNumbersWithIncreasingOrder(n int) []string {
	var res []string
	if n == 1 {
		for i := 0; i < 10; i++ {
			res = append(res, strconv.Itoa(i))
		}

		return res
	}

	nums := 0
	last := 0
	solve(n, &res, &nums, last)
	return res
}

/*
The backtrack step (*nums /= 10) removes the last digit after each full set of numbers starting with a given first digit.
This ensures that when you move to the next i in the outer loop, nums is reset to 0, ready to start building the next set of numbers.
*/
