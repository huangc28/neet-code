package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}})
	log.Printf("debug %v", output)

	output2 := countComponents(4, [][]int{{2, 3}, {1, 2}, {1, 3}})
	log.Printf("debug %v", output2)

	output3 := countComponents(4, [][]int{{0, 1}, {2, 3}, {1, 2}})
	log.Printf("debug %v", output3)
}
