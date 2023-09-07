package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	ans := 0
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
