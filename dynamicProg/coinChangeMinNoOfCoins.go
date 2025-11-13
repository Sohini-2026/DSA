package dynamicProg

/*
Coin Change Problem Minimum Numbers of coins
Given a value V, if we want to make change for V cents, and we have infinite supply of each of C = { C1, C2, .. , Cm} valued coins, what is the minimum number of coins to make the change?
Example:

Input: coins[] = {25, 10, 5}, V = 30
Output: Minimum 2 coins required
We can use one coin of 25 cents and one of 5 cents
*/

import "math"

func CoinChangeMinNoOfCoins(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for j := 1; j <= amount; j++ {
		dp[0][j] = math.MaxInt32 - 1
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= amount; j++ {
			if coins[i-1] <= j {
				dp[i][j] = min(dp[i-1][j], 1+dp[i][j-coins[i-1]])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[n][amount] == math.MaxInt32-1 {
		return -1
	}
	return dp[n][amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
