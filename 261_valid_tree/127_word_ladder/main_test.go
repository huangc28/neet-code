package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := ladderLength2("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	log.Printf("debug output %v", output)
}
