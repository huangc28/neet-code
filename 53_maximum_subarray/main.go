package main

/*
kadane's algorithm

Calculate the maximum subarray sum at each position.
For example: [-2,1,-3,4,-1,2,1,-5,4]

dp[0], what is the maximum subarray sum for subarray [-2], -2
dp[1], what is the maximum subarray sum for subarray [-2,1], max(1, -2+1) = 1
dp[2], what is the maximum subarray sum for subarray [-2,1,-3], max(-3, -3+1) = -2
...
dp[len(nums)-1]

thus, the formula for maxiumum subarray sum at given position would be:

dp[n] = max(nums[n], nums[n]+dp[n-1])
*/
func maxSubArray(nums []int) int {
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i-1]+nums[i])
		maxSum = max(nums[i], maxSum)
	}

	return maxSum

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
