package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	output := canPartition([]int{1, 2, 3, 5, 17, 6, 14, 12, 6})
	assert.Equal(t, true, output)

	output2 := canPartition([]int{1, 5, 11, 5})
	assert.Equal(t, true, output2)

	output3 := canPartition([]int{1, 2, 3, 5})
	assert.Equal(t, false, output3)
}
