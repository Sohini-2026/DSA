package week1

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		if !IsLetter(runes[left]) {
			left++
			continue
		}
		if !IsLetter(runes[right]) {
			right--
			continue
		}

		if !strings.EqualFold(string(runes[left]), string(runes[right])) {
			return false
		}

		left++
		right--
	}

	return true

}

func IsLetter(r rune) bool {
	return unicode.IsLetter(r)
}
