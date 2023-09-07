package main

import "math"

func maxProduct_2(nums []int) int {
	prefMax, prefProd := math.MinInt32, 1

	// (max,prod)
	// (-2,-2), ()
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			prefMax = max(prefMax, nums[i])
			prefProd = 1
			continue
		}

		prefProd *= nums[i]
		prefMax = max(prefMax, prefProd)
	}

	sufMax, sufProd := math.MinInt32, 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			sufMax = max(sufMax, nums[i])
			sufProd = 1
			continue
		}

		sufProd *= nums[i]
		sufMax = max(sufMax, sufProd)
	}

	return max(prefMax, sufMax)
}
