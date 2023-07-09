package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	arr1 := []string{"2", "1", "+", "3", "*"}
	res := evalRPN(arr1)
	log.Printf("debug res %v", res)
}

func TestCase2(t *testing.T) {
	arr2 := []string{"4", "13", "5", "/", "+"}
	res := evalRPN(arr2)
	log.Printf("debug res %v", res)
}
