package main

import (
	"log"
	"testing"
)

/*
1.

	ei
	ab
*/
func TestCase1(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	inPerm := CheckInclusion(s1, s2)
	log.Printf("debug ans %v", inPerm)
}

func TestCase2(t *testing.T) {
	s1 := "ab"
	s2 := "eidboaoo"
	inPerm := CheckInclusion(s1, s2)
	log.Printf("debug ans %v", inPerm)
}
