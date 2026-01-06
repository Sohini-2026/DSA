package week1

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > minPrice {
			maxProfit = max(maxProfit, prices[i]-minPrice)
		} else {
			minPrice = prices[i]
		}
	}

	return maxProfit

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
