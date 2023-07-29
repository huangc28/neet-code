package main

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[i][0]
	})

	stack := make([][]int, len(intervals))
	stack = append(stack, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if stack[len(stack)-1][1] > intervals[i][1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			merged := []int{
				min(top[0], intervals[i][0]),
				max(top[1], intervals[i][1]),
			}
			stack = append(stack, merged)
		} else {
			stack = append(stack, intervals[i])
		}
	}
	return stack
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
