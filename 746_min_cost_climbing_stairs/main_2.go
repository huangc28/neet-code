package main

func minCostClimbingStairs_recurive(cost []int) int {
	dp := make(map[int]int)
	return min(helper(0, cost, dp), helper(1, cost, dp))
}

func helper(p int, cost []int, dp map[int]int) int {
	if v, ok := dp[p]; ok {
		return v
	}

	if p >= len(cost) {
		return 0
	}

	currCost := min(
		helper(p+1, cost, dp), //
		helper(p+2, cost, dp), // 15 + 0
	) + cost[p]

	dp[p] = currCost

	return dp[p]
}
