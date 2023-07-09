package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	bombs := [][]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 4, 2},
		{4, 5, 3},
		{5, 6, 4},
	}

	count := maximumDetonation(bombs)
	log.Printf("debug %v", count)
}

func TestCase2(t *testing.T) {
	bombs := [][]int{
		{1, 1, 5},
		{10, 10, 5},
	}

	count := maximumDetonation(bombs)
	log.Printf("debug ans %v", count)
}

func TestCase3(t *testing.T) {
	bombs := [][]int{{56, 80, 2}, {55, 9, 10}, {32, 75, 2}, {87, 89, 1}, {61, 94, 3}, {43, 82, 9}, {17, 100, 6}, {50, 6, 7}, {9, 66, 7}, {98, 3, 6}, {67, 50, 2}, {79, 39, 5}, {92, 60, 10}, {49, 9, 9}, {42, 32, 10}}
	count := maximumDetonation_2(bombs)
	log.Printf("debug ans %v", count)
}
