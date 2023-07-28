package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	dig1 := '1' - '0'
	dig2 := '0' - '0'

	// ^ xor operator
	res := dig1 ^ dig2
	assert.Equal(t, "1", string(res+'0'))

	dig1 = '1' - '0'
	dig2 = '1' - '0'
	dig3 := '1' - '0'
	res = dig1 ^ dig2 ^ dig3
	assert.Equal(t, "1", string(res+'0'))
}
