package main

func maxProfit(prices []int) int {
	b, s := 0, 0
	maxProfit := 0

	for ; s < len(prices); s++ {
		if prices[b] == prices[s] {
			continue
		}

		// We can not sell a stock at a lost.
		if prices[b] > prices[s] {

		} else {
			// we can sell a stock at a profit
			maxProfit += prices[s] - prices[s]
		}

		// Try buying stock on the sell day we see if we can
		// make profit further on.
		b = s
	}

	return maxProfit
}
