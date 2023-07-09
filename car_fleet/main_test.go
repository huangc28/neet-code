package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}

	fleetCount := carFleet(target, position, speed)
	log.Printf("debug fleetCount %v", fleetCount)
}

func TestCase2(t *testing.T) {
	target := 100
	position := []int{0, 2, 4}
	speed := []int{4, 2, 1}

	fleetCount := carFleet(target, position, speed)
	log.Printf("debug fleetCount %v", fleetCount)
}

func TestCase3(t *testing.T) {
	target := 10
	position := []int{2, 4}
	speed := []int{3, 2}
	fleetCount := carFleet(target, position, speed)
	log.Printf("debug fleetCount %v", fleetCount)
}
