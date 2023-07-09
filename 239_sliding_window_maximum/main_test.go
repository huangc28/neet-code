package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	max := maxSlidingWindow(nums, k)
	log.Printf("debug %v", max)
}
