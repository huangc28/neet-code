package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	sample := "babad"
	odd1 := getOddPal(2, sample)
	assert.Equal(t, "aba", odd1)

	sampleEven := "cbbd"
	even1 := getEvenPal(2, sampleEven)
	log.Printf("debug %v", even1)

	output := longestPalindrome("babad")
	log.Printf("debug output %v", output)
}
