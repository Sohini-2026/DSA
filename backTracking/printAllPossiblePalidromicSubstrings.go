package backtracking

import "fmt"

/*
Given a String S, Find all possible Palindromic partitions of the given String.


Example 1:

Input:
S = "geeks"
Output:
g e e k s
g ee k s
Explanation:
All possible palindromic partitions
are printed.
*/

func PrintAllPossiblePalindromicSubstrings(s string) [][]string {
	var res [][]string
	var path []string
	n := len(s)
	backtrackPalindrome(s, 0, n, path, &res)
	return res
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func backtrackPalindrome(s string, start, n int, path []string, res *[][]string) {
	if start == n {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for end := start; end < n; end++ {
		if isPalindrome(s, start, end) {
			fmt.Println("Found palindrome:", s[start:end+1], "from index", start, "to", end)
			path = append(path, s[start:end+1])
			backtrackPalindrome(s, end+1, n, path, res)
			path = path[:len(path)-1] // backtrack
		}
	}
}
