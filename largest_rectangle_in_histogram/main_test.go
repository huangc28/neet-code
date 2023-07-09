package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	arr1 := []int{2, 1, 5, 6, 2, 3}
	maxArea := largestRectangleArea(arr1)
	log.Printf("max %v", maxArea)
}
