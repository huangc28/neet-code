package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	res := kClosest([][]int{{1, 3}, {-2, 2}}, 1)
	log.Printf("debug %v", res)
}

func TestCase2(t *testing.T) {
	res := kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)
	log.Printf("debug %v", res)
}
