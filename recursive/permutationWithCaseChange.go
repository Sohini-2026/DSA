package recursive

func PermutationWithCaseChange(output, input string) {
	if len(input) == 0 {
		println(output)
		return
	}

	// Include the first character in lowercase
	output1 := output + string(input[0])
	// Include the first character in uppercase
	output2 := output + string(input[0]-32) // ASCII conversion to uppercase
	input = input[1:]                       // Exclude the first character

	PermutationWithCaseChange(output1, input)
	PermutationWithCaseChange(output2, input)
}
