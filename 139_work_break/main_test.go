package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	output := wordBreak_topdown("leetcode", []string{"leet", "code"})
	assert.Equal(t, true, output)

	output2 := wordBreak_topdown("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	assert.Equal(t, false, output2)
}

func TestCase2(t *testing.T) {
	output := wordBreak_bottom_up("leetcode", []string{"leet", "code"})
	assert.Equal(t, true, output)

	output2 := wordBreak_bottom_up("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	assert.Equal(t, false, output2)
}
