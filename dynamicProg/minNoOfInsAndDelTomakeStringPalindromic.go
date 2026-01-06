package dynamicProg

func MinNoOfInsAndDelToMakeStringPalindromic(s string) (int, int) {
	n := len(s)
	s1 := reverseS([]byte(s))

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == s1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	lpsLength := dp[n][n]
	minIns := n - lpsLength
	minDel := n - lpsLength

	return minIns, minDel
}
