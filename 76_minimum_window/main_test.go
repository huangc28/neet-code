package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	s1 := "ADOBECODEBANC"
	t1 := "ABC"

	minWin1 := minWindow(s1, t1)
	log.Printf("debug ans %v", minWin1)
}
