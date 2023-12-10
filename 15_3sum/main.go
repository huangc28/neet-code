package main

import (
	"sort"
)

func ThreeSums(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// move left pointer to right
				k -= 1

				// if we encounter the same number, skip it
				// since we don't want to have duplicate
				// triplet
				for j < k && nums[k] == nums[k+1] {
					k -= 1
				}
			} else if sum > 0 {
				k -= 1
			} else {
				j += 1
			}
		}
	}

	return res
}
