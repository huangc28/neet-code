package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	output := checkValidString("(*)")
	assert.Equal(t, true, output)

	output2 := checkValidString("(*))")
	assert.Equal(t, true, output2)
}
