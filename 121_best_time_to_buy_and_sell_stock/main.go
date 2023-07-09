package main

func maxProfit(prices []int) int {
	b := 0
	s := 1
	maxProfit := 0

	for s < len(prices) {
		if prices[s] <= prices[b] {
			b = s
			s = b + 1
		} else {
			maxProfit = max(maxProfit, prices[s]-prices[b])
			s++
		}
	}

	return maxProfit
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
