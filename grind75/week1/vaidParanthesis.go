package week1

func ValidParentheses1(s string) bool {
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)

	for _, char := range s {
		if open, exists := mapping[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != open { // starts with closing bracket or mismatch
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
