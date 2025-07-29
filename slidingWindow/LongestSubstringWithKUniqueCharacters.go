package slidingwindow

// Pick a toy from the shelf and put it in the basket. If the basket has more than k toys, remove the toy from the leftmost side of the basket.
// Return the longest substring with exactly k unique characters.
// If there are no such substrings, return an empty string.
// Time Complexity: O(n), where n is the length of the string.
// Space Complexity: O(k), where k is the number of unique characters in the string.
// Example: For the string "aabacbebebe" and k = 3,
// the longest substring with exactly 3 unique characters is "cbebebe".
func LongestSubstringWithKUniqueCharacters(s string, k int) string {
	if len(s) == 0 || k <= 0 {
		return "" // Invalid input
	}

	i, j := 0, 0
	charCount := make(map[rune]int)
	maxLen := 0
	startIdx := -1

	for j < len(s) {
		charCount[rune(s[j])]++ // Add the current character to the count

		for len(charCount) > k { // Shrink the window from the left if we have more than k unique characters
			charCount[rune(s[i])]--
			if charCount[rune(s[i])] == 0 {
				delete(charCount, rune(s[i]))
			}
			i++
		}

		if (j - i + 1) > maxLen { // Check if we found a larger window
			maxLen = j - i + 1
			startIdx = i
		}
		j++ // Move to the next character
	}

	if startIdx != -1 {
		return s[startIdx : startIdx+maxLen]
	}

	return "" // No valid substring found
}
