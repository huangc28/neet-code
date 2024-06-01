package main

import "log"

func findDuplicate(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	log.Printf("debug %v %v", slow, fast)

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	log.Printf("debug  2 %v %v", slow, fast)

	return slow
}
