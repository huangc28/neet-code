package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 7, 4, 1, 8, 1}
	output := lastStoneWeight(nums)
	log.Printf("debug %v", output)
}
