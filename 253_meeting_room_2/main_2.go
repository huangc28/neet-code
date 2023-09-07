package main

import (
	"log"
	"sort"
)

type TimeInfo struct {
	Time    int
	IsStart bool
}

func minMeetingRooms_2(intervals [][]int) int {
	starts := make([]TimeInfo, 0, len(intervals))
	ends := make([]TimeInfo, 0, len(intervals))

	for _, interval := range intervals {
		starts = append(starts, TimeInfo{
			Time:    interval[0],
			IsStart: true,
		})
		ends = append(ends, TimeInfo{
			Time:    interval[1],
			IsStart: false,
		})
	}

	times := append(starts, ends...)
	sort.Slice(times, func(i, j int) bool {
		if times[i].Time == times[j].Time {
			return !times[i].IsStart
		}
		return times[i].Time < times[j].Time
	})

	log.Printf("debug %v", times)

	maxRooms := 0
	currRooms := 0

	for _, time := range times {
		if time.IsStart {
			currRooms += 1
		} else {
			currRooms -= 1
		}
		maxRooms = max(maxRooms, currRooms)
	}

	return maxRooms
}
