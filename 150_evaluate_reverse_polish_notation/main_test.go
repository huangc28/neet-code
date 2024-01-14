package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	res := evalRPN(tokens)
	assert.Equal(t, 9, res)

	tokens2 := []string{"4", "13", "5", "/", "+"}
	res2 := evalRPN(tokens2)
	assert.Equal(t, 6, res2)
}
