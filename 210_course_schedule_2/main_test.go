package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	log.Printf("debug %v", output)
}
