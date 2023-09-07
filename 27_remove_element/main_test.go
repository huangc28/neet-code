package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{3, 2, 2, 3}
	removeElement(input, 3)
	log.Printf("debug %v", input)

	input2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(input2, 2)
	log.Printf("debug %v", input2)
}
