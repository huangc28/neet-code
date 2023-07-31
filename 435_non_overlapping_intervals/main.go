package main

import (
	"sort"
)

/*
if we found 2 intervals that are overlapping
remove the interval that ends later. for example

(1, 2) and (1, 3) are overlapping, we remove (1,3) since it ends later
*/

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	stack := make([][]int, 0)
	stack = append(stack, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if stack[len(stack)-1][1] > intervals[i][0] {
			// overlapped, check who ends earlier
			if stack[len(stack)-1][1] < intervals[i][1] {
				// don't do anything
			} else {
				// current interval ends earlier, update
				// the end time
				stack[len(stack)-1][1] = intervals[i][1]
			}

		} else {
			stack = append(stack, intervals[i])
		}
	}

	return len(intervals) - len(stack)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
