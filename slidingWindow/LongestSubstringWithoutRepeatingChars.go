package slidingwindow

func LongestSubstringWithoutRepeatingCharacters(s string) string {
	if len(s) == 0 {
		return "" // Invalid input
	}

	i, j := 0, 0
	charCount := make(map[rune]int)
	maxLen := 0
	startIdx := -1

	for j < len(s) {
		charCount[rune(s[j])]++ // Add the current character to the count

		for charCount[rune(s[j])] > 1 { // Shrink the window from the left if we have a repeating character
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
