package backtracking

/*
Given a string s and a dictionary of words dict of length n, add spaces in s to construct a sentence where each word is a valid dictionary word. Each dictionary word can be used more than once. Return all such possible sentences.

Follow examples for better understanding.

Example 1:

Input: s = "catsanddog", n = 5
dict = {"cats", "cat", "and", "sand", "dog"}
Output: (cats and dog)(cat sand dog)
Explanation: All the words in the given
sentences are present in the dictionary
*/

func WordBreak(s string, dict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range dict {
		wordSet[word] = true
	}

	var res []string
	var path []string
	backtrackWordBreak(s, 0, wordSet, path, &res)
	return res
}

func backtrackWordBreak(s string, start int, wordSet map[string]bool, path []string, res *[]string) {
	if start == len(s) {
		sentence := ""
		for i, word := range path {
			if i > 0 {
				sentence += " "
			}
			sentence += word
		}
		*res = append(*res, sentence)
		return
	}

	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		if wordSet[word] {
			path = append(path, word)
			backtrackWordBreak(s, end, wordSet, path, res)
			path = path[:len(path)-1] // backtrack
		}
	}
}
