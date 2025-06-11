package recursive

import (
	"strings"
	"unicode"
)

func LetterCasePermutation(output, input string) {
	if len(input) == 0 {
		println(output)
		return
	}

	// Include the first character in lowercase
	if !unicode.IsDigit(rune(input[0])) {
		output1 := output + strings.ToLower(string(input[0]))
		output2 := output + strings.ToUpper(string(input[0]))
		input = input[1:] // Exclude the first character
		LetterCasePermutation(output1, input)
		LetterCasePermutation(output2, input)
	} else {
		// If the character is a digit, just append it
		output1 := output + string(input[0])
		input = input[1:] // Exclude the first character
		LetterCasePermutation(output1, input)
	}

}
