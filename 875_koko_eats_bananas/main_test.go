package main

import (
	"log"
	"testing"
)

func TestSol(t *testing.T) {
	sol := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	log.Printf("debug sol %v", sol)
	// sol := calcTimeToFinish([]int{3, 6, 7, 11}, 4)
	// log.Printf("debug sol %v", sol)
}
