package recursive

func CreateBalancedParanthesis(output, input string, open, close int) {
	if open == 0 && close == 0 {
		println(output)
		return
	}

	if open > 0 {
		CreateBalancedParanthesis(output+"(", input, open-1, close)
	}

	if close > open {
		CreateBalancedParanthesis(output+")", input, open, close-1)
	}
}
