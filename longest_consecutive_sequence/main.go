package main

import "sort"

/*
[100,4,200,1,3,2]

這些數字中，哪些數字是沒有 left neighbor? 這些數字就是 start of the sequence.
然後 iterate 這串 array numbr by numbr 檢查哪一些數字是 start of the sequence。
如果是 start of the sequence 我們就建立一個新的 sequence.
*/
func longestConsecutive(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	currStreak, maxStreak := 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			currStreak++
		} else if nums[i] == nums[i-1] {
			// ignore
		} else {
			maxStreak = max(currStreak, maxStreak)
			currStreak = 1
		}
	}

	// ..., 101, 102
	return max(maxStreak, currStreak)
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
