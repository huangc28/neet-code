package main

import "fmt"

/*

[1,5,11,5]

sum = 22
target = 22/2 = 11

find a subset that sums up to 11
	    f(0,11)
	     /        \
	f(1,10)       f(1, 11)
	  /  \           /   \
      f(2,5) f(2,10)  f(2,5) f(2,11)
      true     /
              /
	     f(3,-1)



*/

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	dp := make(map[string]bool)
	target := sum / 2
	return helper(len(nums)-1, target, nums, dp)
}

func helper(idx, target int, nums []int, dp map[string]bool) bool {
	if target == 0 {
		return true
	} else if idx == 0 {
		return nums[0] == target
	}

	if v, ok := dp[fmt.Sprintf("%d-%d", idx, target)]; ok {
		return v
	}

	notTake := helper(idx-1, target, nums, dp)
	take := false
	if target >= nums[idx] {
		take = helper(idx-1, target-nums[idx], nums, dp)
	}

	dp[fmt.Sprintf("%d-%d", idx, target)] = take || notTake

	return dp[fmt.Sprintf("%d-%d", idx, target)]
}
