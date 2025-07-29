package slidingwindow

func MinimumWindowSubstring(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return "" // Invalid input
	}

	i, j := 0, 0
	charCount := make(map[rune]int)
	for _, char := range t {
		charCount[char]++
	}

	requiredChars := len(charCount)
	formedChars := 0
	minLen := len(s) + 1
	startIdx := -1

	for j < len(s) {
		charCount[rune(s[j])]--

		if charCount[rune(s[j])] == 0 {
			formedChars++
		}

		for i <= j && formedChars == requiredChars {
			if (j - i + 1) < minLen {
				minLen = j - i + 1
				startIdx = i
			}

			charCount[rune(s[i])]++        // Remove the character at the start of the window
			if charCount[rune(s[i])] > 0 { // If the character count goes above zero, we no longer have a valid window
				formedChars-- //This is to ensure we only count characters that are still required
			}
			i++
		}
		j++
	}

	if startIdx == -1 {
		return "" // No valid substring found
	}
	return s[startIdx : startIdx+minLen]
}
