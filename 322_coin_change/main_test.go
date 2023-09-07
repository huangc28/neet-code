package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	output := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(t, 3, output)

	output2 := coinChange([]int{2}, 3)
	assert.Equal(t, -1, output2)
}
