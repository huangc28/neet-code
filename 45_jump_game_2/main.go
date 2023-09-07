package main

import "math"

func jump(nums []int) int {
	// what is the minimum step that I should take to reach last index?
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ { // O(n)
		dp[i] = math.MaxInt32
	}
	dp[len(dp)-1] = 0 // 0 steps to reach the last index from the last index

	for i := len(nums) - 2; i >= 0; i-- { // O(n)
		// among the jumps that I can take, what is the minimum steps to reach the top?
		for jumps := 1; jumps <= nums[i]; jumps++ { // O(m)
			// 1
			// ---
			// 1+dp[2+1]
			// ---
			// 1+dp[1+1] 3
			// 1+dp[1+2] 2
			// 1+dp[1+3] 1
			// ---
			// 1+dp[0+1] 2
			// 1+dp[0+2] 3
			if i+jumps < len(dp) {
				dp[i] = min(dp[i], 1+dp[i+jumps])
			}
		}
	}

	return dp[0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
