package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	res := canFinish(2, [][]int{{1, 0}})
	assert.Equal(t, true, res)

	res2 := canFinish(2, [][]int{{1, 0}, {0, 1}})
	assert.Equal(t, false, res2)
}
