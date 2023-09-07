package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input, 3)
	log.Printf("debug %v", input)
}

func TestCase2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input, 3)
	log.Printf("debug %v", input)
}
