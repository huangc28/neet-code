package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
	log.Printf("debug output %v", output)
}
