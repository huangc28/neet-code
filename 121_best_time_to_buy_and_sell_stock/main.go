package main

func maxProfit(prices []int) int {
	profit := 0

	i := 0 // buying on current day

	for i < len(prices)-1 { // 1 more future day to compare with
		j := i + 1 // j is future day

		// future day's price is lower, I might as well
		// buy on future day
		if prices[j] <= prices[i] {
			i = j
			continue
		}

		// future day's price is higher, compare the max
		// profix to aquire.
		for j < len(prices) && prices[j] > prices[i] {
			profit = max(profit, prices[j]-prices[i])
			j += 1
		}

		// Proceed buying day forward
		i = j
	}

	return profit
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
