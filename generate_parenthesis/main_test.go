package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	n := 3
	res := generateParenthesis(n)
	log.Printf("debug %v", res)
}
