package week1

// length check + hash map count
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	for _, char := range t {
		counts[char]--
		if counts[char] < 0 { // more occurrences in t than in s
			return false
		}
	}

	return true
}
