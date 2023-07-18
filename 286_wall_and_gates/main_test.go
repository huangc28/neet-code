package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	board := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}

	wallsAndGates(board)
	log.Printf("debug %v", board)
}
