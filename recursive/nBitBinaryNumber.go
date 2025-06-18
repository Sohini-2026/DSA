package recursive

func NBitBinaryNumber(n int, output string, ones, zeros int) {
	if n == 0 {
		println(output)
		return
	}
	// Always allowed to add '1'
	NBitBinaryNumber(n-1, output+"1", ones+1, zeros)
	// Only add '0' if ones > zeros
	if ones > zeros {
		NBitBinaryNumber(n-1, output+"0", ones, zeros+1)
	}
}

/*
Explanation:
As it requires a binary number of n bits, we can always add '1' to the output.
We can only add '0' if the number of '1's is greater than the number of '0's.
This ensures that the binary number is valid and maintains the condition of having more '1's than '0's at any point.
// This function generates all n-bit binary numbers with the condition that
// the number of '1's is always greater than the number of '0's.
// It uses recursion to build the binary number by adding '1's and '0's
// while maintaining the condition.
// The function takes the number of bits (n), the current output string,
// and counters for the number of '1's and '0's as parameters.
// It prints all valid n-bit binary numbers that satisfy the condition.
// Example usage:
// NBitBinaryNumber(3, "", 0, 0)
// Output:
// 111
// 110
// 101


*/
