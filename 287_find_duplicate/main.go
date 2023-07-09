package main

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

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
