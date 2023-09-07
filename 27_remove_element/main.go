package main

import (
	"math"
	"sort"
)

func removeElement(nums []int, val int) int {
	count := 0
	for idx, num := range nums {
		if num == val {
			nums[idx] = math.MaxInt32
			continue
		}
		count++
	}

	sort.Ints(nums)

	return count
}

func removeElement_2(nums []int, val int) int {
	// index to where we will store num that is different than val
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
