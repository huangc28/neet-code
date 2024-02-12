package main

import (
	"log"
	"testing"
)

func TestShiftZeroRight(t *testing.T) {
	res := (-1) >> 1
	log.Printf("debug %v", res)
}
