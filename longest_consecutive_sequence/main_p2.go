package main

import "sort"

func longestConsecutiveP2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	sortNums(nums)

	maxConsec := 1
	localConsec := 1

	for i := 0; i < len(nums); i++ {
		if i == 0 || abs(nums[i], nums[i-1]) > 1 {
			localConsec = 1
		} else if abs(nums[i], nums[i-1]) == 0 {
			// Ignore
		} else {
			localConsec += 1
			if localConsec > maxConsec {
				maxConsec = localConsec
			}
		}
	}

	return maxConsec
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func sortNums(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}
