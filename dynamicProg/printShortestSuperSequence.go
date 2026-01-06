package dynamicProg

/*
Given two strings X and Y, print the shortest string that has both X and Y as subsequences. If multiple shortest supersequence exists, print any one of them.

Examples:

Input: X = "AGGTAB",  Y = "GXTXAYB"
Output: "AGXGTXAYB" OR "AGGXTXAYB"
OR Any string that represents shortest
supersequence of X and Y

Input: X = "HELLO",  Y = "GEEK"
Output: "GEHEKLLO" OR "GHEEKLLO"
OR Any string that represents shortest
supersequence of X and Y
*/
func PrintShortestSuperSequence(X, Y string) string {
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

	//length of shortest supersequence = m + n - length of LCS

	index := dp[m][n]
	i, j := m, n
	superSeq := make([]byte, m+n-index)
	k := 0

	for i > 0 && j > 0 {
		if X[i-1] == Y[j-1] {
			superSeq[k] = X[i-1]
			i--
			j--
			k++
		} else if dp[i-1][j] > dp[i][j-1] {
			superSeq[k] = X[i-1]
			i--
			k++
		} else {
			superSeq[k] = Y[j-1]
			j--
			k++
		}
	}

	for i > 0 {
		superSeq[k] = X[i-1]
		i--
		k++
	}

	for j > 0 {
		superSeq[k] = Y[j-1]
		j--
		k++
	}

	reverse(superSeq)

	return string(superSeq)
}

func reverse(arr []byte) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
