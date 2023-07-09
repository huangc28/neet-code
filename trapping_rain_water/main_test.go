package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	sampleArr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	totalArea := trap(sampleArr)
	log.Printf("debug %v", totalArea)
}

func TestCase2(t *testing.T) {
	sampleArr := []int{4, 2, 0, 3, 2, 5}
	totalArea := trap(sampleArr)
	log.Printf("debug %v", totalArea)
}
