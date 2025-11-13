package dynamicProg

/*
Printing Longest Common Subsequence
Given two sequences, print the longest subsequence present in both of them.
Example:
LCS for input Sequences “ABCDGH” and “AEDFHR” is “ADH” of length 3.
*/

func PrintLongestCommonSubsequence(X, Y string) string {
	m := len(X)
	n := len(Y)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	index := dp[m][n]
	lcs := make([]byte, index)
	i, j := m, n
	for i > 0 && j > 0 {
		if X[i-1] == Y[j-1] {
			lcs[index-1] = X[i-1]
			i--
			j--
			index--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(lcs)
}
