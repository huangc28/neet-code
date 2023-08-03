package main

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] = cost[i] + min(cost[i+1], cost[i-1])
	}

	return min(cost[0], cost[1])
}

func minCostClimbingStairs2(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] = cost[i] + min(cost[i-1], cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
