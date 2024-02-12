package main

import (
	"math"
)

/*
[3,6,7,11]
[0,1,2,3,4,5,6,7,8,9,10,11], h=8

	0                      12
	      l
	          r
	(0+/12)/2=6
	6 per hour: 1+1+1+1=4, 4 < 8, 應該還可以再更慢
	(0+5)/2=2
	2+3+4+6=15, 15>8 要再更快
	(3+5)/2=4
	1+2+2+3=8, 8=8 剛剛好

	[30,11,23,4,20]
*/
func minEatingSpeed(piles []int, h int) int {
	// Find the largest pile which is maximum of banana that coco can eat
	maxCanEat := piles[0]
	for i := 0; i < len(piles); i++ {
		maxCanEat = max(maxCanEat, piles[i])
	}

	// Perform binary search to approach the speed that we want for koko to
	// eat banana within h hours.
	res := maxCanEat // number of banana koko can eat in 1 hour
	left, right := 0, maxCanEat
	for left <= right {
		mid := (left + right) / 2

		// Calculate the time to finish
		timeToFinish := calcTimeToFinish(piles, mid)

		if timeToFinish <= h { // koko eats too fast try slower
			res = min(res, mid)
			right = mid - 1
		} else { // koko should eat faster
			left = mid + 1
		}
	}

	return res
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

func calcTimeToFinish(piles []int, speed int) int {
	total := 0
	for i := 0; i < len(piles); i++ {
		sf := float64(piles[i]) / float64(speed)
		sfc := math.Ceil(sf)
		total += int(sfc)
	}
	return total
}
