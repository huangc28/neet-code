package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	str := "ABAB"
	k := 2
	maxCount := characterReplacement(str, k)
	assert.Equal(t, maxCount, 4)

	str2 := "AABABBA"
	k2 := 1
	maxCount2 := characterReplacement(str2, k2)
	assert.Equal(t, maxCount2, 4)

	str3 := "ABCDE"
	k3 := 1
	maxCount3 := characterReplacement(str3, k3)
	assert.Equal(t, maxCount3, 2)
}
