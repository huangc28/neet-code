package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	redundant := findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})
	log.Printf("debug %v", redundant)
}
