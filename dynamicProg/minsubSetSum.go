package dynamicProg

/*
Given a set of integers, the task is to divide it into two sets S1 and S2 such that the absolute difference between their sums is minimum.
If there is a set S with n elements, then if we assume Subset1 has m elements, Subset2 must have n-m elements and the value of abs(sum(Subset1) – sum(Subset2)) should be minimum.

Example:
Input:  arr[] = {1, 6, 11, 5}
Output: 1
Explanation:
Subset1 = {1, 5, 6}, sum of Subset1 = 12
Subset2 = {11}, sum of Subset2 = 11
*/

func MinSubsetSumDifference(arr []int) int {
	totalSum := 0
	for _, num := range arr {
		totalSum += num
	}

	n := len(arr)
	target := totalSum / 2

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for s1 := target; s1 >= 0; s1-- {
		if dp[n][s1] {
			s2 := totalSum - s1
			return abs(s2 - s1)
		}
	}

	return totalSum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
