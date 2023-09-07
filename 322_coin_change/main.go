package main

import (
	"math"
)

/*
coins = [1,2,5]
dp = []
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for pa := 1; pa < len(dp); pa++ {
		for _, coin := range coins {
			if coin <= pa {
				dp[pa] = min(1+dp[pa-coin], dp[pa])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
