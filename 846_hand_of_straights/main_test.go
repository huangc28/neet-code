package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := isNStraightHand_2([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3)
	log.Printf("debug %v", output)
}
