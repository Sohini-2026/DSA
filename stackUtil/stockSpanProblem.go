package stackUtil

import (
	"github.com/Sohini-2026/DSA/baseDS"
)

/*
The stock span problem is a financial problem where we have a series of n daily price quotes for a stock and we need to calculate the span of stock’s price for all n days. The span Si of the stock’s price on a given day i is defined as the maximum number of consecutive days just before the given day, for which the price of the stock on the current day is less than or equal to its price on the given day. For example, if an array of 7 days prices is given as {100, 80, 60, 70, 60, 75, 85}, then the span values for corresponding 7 days are {1, 1, 1, 2, 1, 4, 6}.
Example:
Input: Price[] = {100, 80, 60, 70, 60, 75, 85}
Output: Span[] = {1, 1, 1, 2, 1, 4, 6}
*/
func StockSpan(prices []int) []int {
	n := len(prices)
	span := make([]int, n)
	stack := &baseDS.Stack{}

	for i := 0; i < n; i++ {
		// Pop elements from the stack while the stack is not empty and the current price is greater than or equal to the price at the index stored at the top of the stack
		for stack.Size() > 0 && prices[stack.Top()] <= prices[i] {
			stack.Pop()
		}

		// If the stack is empty, it means there are no previous days with a higher price
		if stack.Size() == 0 {
			span[i] = i + 1
		} else {
			// The span is the difference between the current index and the index of the last higher price
			span[i] = i - stack.Top()
		}

		// Push the current index onto the stack
		stack.Push(i)
	}

	return span
}
