package main

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	// calculate prefix
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	// calculate suffix at each position
	prevSuffix := 1
	for i := len(nums) - 2; i >= 0; i-- {
		currSuffix := nums[i+1] * prevSuffix
		ans[i] = ans[i] * currSuffix
		prevSuffix = currSuffix
	}

	return ans
}
