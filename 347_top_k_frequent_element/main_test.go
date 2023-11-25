package main

import (
	"log"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	sol := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	log.Printf("debug %v", sol)
}
