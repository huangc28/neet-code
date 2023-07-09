package main

import (
	"log"
	"math"
)

func maximumDetonation(bombs [][]int) int {
	maxDenotate := 0

	for i := 0; i < len(bombs); i++ {
		bombToStart := bombs[i]

		// create a new slice
		bombsToTest := append([][]int{}, bombs[:i]...)
		bombsToTest = append(bombsToTest, bombs[i+1:]...)
		count := maximumDetonationHelper(bombToStart, bombsToTest)
		maxDenotate = max(maxDenotate, count)
	}

	return maxDenotate
}

func maximumDetonationHelper(bombToStart []int, bombsToTest [][]int) int {
	// queue of bombs
	queue := make([][]int, 0)
	queue = append(queue, bombToStart)
	detonateCount := 0

	for len(queue) > 0 {
		bombToStart := queue[0]
		queue = queue[1:]
		detonateCount++
		bombsCannotBeDetonated := make([][]int, 0)

		for i := 0; i < len(bombsToTest); i++ {
			log.Printf("debug 1 %v", len(bombsToTest))
			bombToTest := bombsToTest[i]
			if canDetonate(bombToStart, bombToTest) {
				// remove bombToTest from bombsToTest
				// add bombToTest to queue
				queue = append(queue, bombToTest)
			} else {
				// skip testing
				bombsCannotBeDetonated = append(bombsCannotBeDetonated, bombsToTest[i])
			}
		}
		bombsToTest = bombsCannotBeDetonated
	}

	return detonateCount
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func canDetonate(bombToStart, bombToTest []int) bool {
	xDis := math.Pow(float64(bombToStart[0])-float64(bombToTest[0]), 2)
	yDis := math.Pow(float64(bombToStart[1])-float64(bombToTest[1]), 2)
	dis := math.Sqrt(xDis + yDis)

	return float64(bombToStart[2]) >= dis
}
