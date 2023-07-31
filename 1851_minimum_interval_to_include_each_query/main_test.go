package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5})
	log.Printf("debug %v", output)
}

func TestCase2(t *testing.T) {
	output := minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22})
	log.Printf("debug %v", output)
}
