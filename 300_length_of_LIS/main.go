package main

func lengthOfLIS_topdown(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			// 當前的數字一定要比後面的數字小，才是 increasing sequence
			if nums[i] < nums[j] {
				// 算出從當前步數走，最長的 subsequence
				// 當前 index 比較長，還是我加上前一個 index 最長的 subsequence?
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

/*
[1,2,4,3]

1->2
  2->4
  2->3
1->4
1->3
2->4
2->3

Time: O(n^2)
*/

func lengthOfLIS_bottomup(nums []int) int {
	dp := make(map[int]int)
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if v, ok := dp[idx]; ok {
			return v
		}

		// base case at the current
		dp[idx] = 1

		for j := idx + 1; j < len(nums); j++ {
			if nums[idx] < nums[j] {
				dp[idx] = max(dp[idx], 1+dfs(j))
			}
		}

		return dp[idx]
	}

	ans := 0
	for i := 0; i < len(nums); i++ { // O(n)
		ans = max(ans, dfs(i))
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
