package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	sol := solveNQueens(4)
	log.Printf("debug %v", sol)
}
