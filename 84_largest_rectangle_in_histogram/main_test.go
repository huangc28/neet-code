package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	ans := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	assert.Equal(t, ans, 10)

	ans2 := largestRectangleArea([]int{2, 1, 2})
	assert.Equal(t, 3, ans2)
}
