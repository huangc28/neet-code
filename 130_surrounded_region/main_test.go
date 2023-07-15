package main

import (
	"log"
	"testing"
)

var board = [][]byte{
	{'X', 'X', 'X', 'X'},
	{'X', '0', '0', 'X'},
	{'X', 'X', '0', 'X'},
	{'X', '0', 'X', 'X'},
}

func TestCase1(t *testing.T) {
	solve(board)
	log.Printf("debug %v", board)
}
