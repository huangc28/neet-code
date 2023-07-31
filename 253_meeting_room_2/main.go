package main

import "sort"

/*
if 2 intervals overlapped, it means we need an extra conference room.

[0, 30] 1 room
[5, 10] 1 room [15, 20] would use this room, it end first


if 2 intervals not overlapped

[0, 30][31, 32] we need only 1 room, the lastEndTime would be 32

[[2,15],[4,9],[9,29],[16,23],[36,45]]
                                i
stack = [(4,29),(2,45)]

[2,15] [4,9] overlapped room++ 9 ends first
[4,9] [9,29] no overlapped, reuse the same room, update endtime, (2,15) ends first swap
[2,15][16,23] no overlapped, reuse the same room, update endtime, (2,23) ends first, stay
[2,23][36,45] no overlapped, reuse the same room, update end time, (4,29) ends first, swap

[[1080,2236],[3325,5014],[1397,3851],[54,1519],[1794,2238],[2907,4858]]

[[54,1519],[1080,2236],[1397,3851],[1794.2238],[2907,4858],[3325,5014]]
                i           i

[[1080,2236],[1397, 3851], [54,1519]]
[54,1519],[1080,2236] overlapped
[54,1519][1397,3851] overlapped

*/

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	if len(intervals) == 1 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	stack := make([][]int, 0)
	stack = append(stack, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if stack[len(stack)-1][1] > intervals[i][0] {
			// we need extra room
			stack = append(stack, intervals[i])

			// which room ends first?
			for j := len(stack) - 1; j > 0; j-- {
				if stack[j-1][1] < stack[j][1] {
					swap(stack, j-1, j)
				}
			}
		} else {
			// no overlapping, we can reuse the room, simply update the time
			stack[len(stack)-1][1] = max(
				stack[len(stack)-1][1],
				intervals[i][1],
			)

			// which room ends first?
			for j := len(stack) - 1; j > 0; j-- {
				if stack[j-1][1] < stack[j][1] {
					swap(stack, j-1, j)
				}
			}
		}
	}

	return len(stack)
}

func swap(nums [][]int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
