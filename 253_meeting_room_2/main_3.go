package main

import "sort"

func minMeetingRooms_3(intervals [][]int) int {
	starts := make([]int, 0, len(intervals))
	ends := make([]int, 0, len(intervals))

	for _, interval := range intervals {
		starts = append(starts, interval[0])
		ends = append(ends, interval[1])
	}

	sort.Ints(starts)
	sort.Ints(ends)

	sp, ep := 0, 0
	currRooms, maxRooms := 0, 0

	for sp < len(starts) && ep < len(ends) {
		if starts[sp] == ends[ep] || starts[sp] > ends[ep] {
			ep++
			currRooms--
		} else { // a meeting is starting
			sp++
			currRooms++
		}
		maxRooms = max(maxRooms, currRooms)
	}

	for sp < len(starts) {
		currRooms++
		sp++
	}

	for ep < len(ends) {
		currRooms--
		ep++
	}

	return maxRooms
}
