package main

import (
	"log"
	"testing"
)

func Test3Sum(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	res := ThreeSums(arr)
	log.Printf("debug 1 %v", res)
}
