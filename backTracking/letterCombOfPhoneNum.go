package backtracking

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.


Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
*/

var digitToLetters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinationsOfPhoneNumber(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var res []string
	var path []byte
	backtrackLetterCombinations(digits, 0, path, &res)
	return res
}

func backtrackLetterCombinations(digits string, index int, path []byte, res *[]string) {
	if index == len(digits) {
		*res = append(*res, string(path))
		return
	}

	letters := digitToLetters[digits[index]]
	for i := 0; i < len(letters); i++ {
		path = append(path, letters[i])
		backtrackLetterCombinations(digits, index+1, path, res)
		path = path[:len(path)-1] // backtrack
	}
}
