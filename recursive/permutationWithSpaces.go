package recursive

func PermutationWithSpaces(output, input string) {
	if len(input) == 0 {
		println(output)
		return
	}

	// Include the first character without space
	output1 := output + " " + string(input[0])
	// Include the first character with space
	output2 := output + string(input[0])
	input = input[1:] // Exclude the first character
	PermutationWithSpaces(output1, input)

	PermutationWithSpaces(output2, input)

}
