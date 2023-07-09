package main

import "sort"

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	left, right := 1, piles[len(piles)-1]
	res := right

	for left <= right {
		midSpeed := (left + right) / 2
		timeToFinish := calcTimeToFinish(midSpeed, piles)

		// eating faster than
		if timeToFinish <= h {
			res = min(res, timeToFinish)
			right = midSpeed - 1
		} else {
			left = midSpeed + 1

		}
	}

	return res
}


func calcTimeToFinish(speed int, piles []int) int {
	for i := 0; i < len(piles) i++ {
	}
	return 0
}
