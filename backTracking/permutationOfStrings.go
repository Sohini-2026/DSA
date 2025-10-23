package backtracking

func PermuteStrings(strs []string) [][]string {
	var res [][]string
	backtrack(0, strs, &res)

	return res
}

func backtrack(start int, strs []string, res *[][]string) {
	if start == len(strs) {
		tmp := make([]string, len(strs))
		copy(tmp, strs)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(strs); i++ {
		strs[start], strs[i] = strs[i], strs[start]
		backtrack(start+1, strs, res)
		strs[start], strs[i] = strs[i], strs[start]
	}
}
