package dynamicProg

/*
Given a rod of length n inches and an array of prices that contains prices of all pieces of size smaller than n. Determine the  maximum value obtainable by cutting up the rod and selling the pieces.
Example:
if length of the rod is 8 and the values of different pieces are given as following, then the maximum obtainable value is 22 (by cutting in two pieces of lengths 2 and 6)


length   | 1   2   3   4   5   6   7   8
--------------------------------------------
price    | 1   5   8   9  10  17  17  20
*/

func RodCuttingProblem(prices []int, n int) int {
	np := len(prices)
	dp := make([][]int, np+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= np; j++ {
			if i <= j {
				dp[i][j] = max(dp[i-1][j], prices[i-1]+dp[i][j-i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[np][n]

}

/*
i = 1 (piece length 1, price 1)
For all j from 1 to 8:

You can only use piece of length 1.
dp[1][j] = max(dp[0][j], prices[0] + dp[1][j-1])
This is just j * 1.
j	dp[1][j]
1	1
2	2
3	3
4	4
5	5
6	6
7	7
8	8

i = 2 (piece length 2, price 5)
For each j:

If j < 2, can't use piece 2: just copy above.
If j >= 2, can use piece 2:
dp[2][j] = max(dp[1][j], prices[1] + dp[2][j-2])
j	dp[2][j]	Calculation
1	1	can't use piece 2
2	5	max(2, 5 + 0) = 5
3	6	max(3, 5 + 1) = 6
4	10	max(4, 5 + 5) = 10
5	11	max(5, 5 + 6) = 11
6	15	max(6, 5 + 10) = 15
7	16	max(7, 5 + 11) = 16
8	20	max(8, 5 + 15) = 20

i = 3 (piece length 3, price 8)
For each j:

If j < 3, can't use piece 3: just copy above.
If j >= 3, can use piece 3:
dp[3][j] = max(dp[2][j], prices[2] + dp[3][j-3])
j	dp[3][j]	Calculation
1	1	can't use piece 3
2	5	can't use piece 3
3	8	max(6, 8 + 0) = 8
4	10	max(10, 8 + 1) = 10
5	13	max(11, 8 + 5) = 13
6	16	max(15, 8 + 8) = 16
7	18	max(16, 8 + 10) = 18
8	21	max(20, 8 + 13) = 21

Continuing this way up to i = 8, we fill the entire dp table.
At the end, dp[8][8] will give the maximum obtainable value for a rod of length 8.
*/

/*
Absolutely! Let's clarify where repetitive use happens in your code using a small example.

Key Line for Repetitive Use
prices[i-1] + dp[i][j-i] means:
You cut off a piece of length i (get price prices[i-1])
You still have j-i length left, and you can use pieces up to length i again (including i itself!)
This is what allows repetitive use of the same piece length.

Example: Rod Length 5, Prices = [1, 5, 8, 9, 10]
Let's see how you can use piece length 2 (price 5) multiple times:

Suppose i = 2 (piece length 2), j = 4 (rod length 4):
Option 1: Don't use piece 2 → dp[1][4] (only pieces up to length 1)
Option 2: Use piece 2:
Cut off length 2 (price 5)
Remaining rod: 2
Best price for remaining rod: dp[2][2] (can use piece 2 again!)
So:

Here, you used piece length 2 twice (two cuts of length 2, total price 10).

Generalization
Whenever you use prices[i-1] + dp[i][j-i],
you are allowed to use piece length i again for the remaining rod.
This is the unbounded knapsack pattern, allowing repetitive use.
Summary
The repetitive use happens in the recursive part: dp[i][j-i] (you can use the same piece again for the leftover rod).
The DP table builds up solutions by considering all ways to cut the rod, including using the same piece length multiple times.
Let me know if you want a step-by-step trace for a specific rod length and piece!
*/
