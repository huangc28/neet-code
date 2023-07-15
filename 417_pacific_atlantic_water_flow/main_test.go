package main

import (
	"log"
	"testing"
)

var map1 = [][]int{
	{1, 2, 2, 3, 5},
	{3, 2, 3, 4, 4},
	{2, 4, 5, 3, 1},
	{6, 7, 1, 4, 5},
	{6, 1, 1, 2, 4},
}

func TestCase1(t *testing.T) {
	output := pacificAtlantic(map1)
	log.Printf("debug 2 %v", output)
}
