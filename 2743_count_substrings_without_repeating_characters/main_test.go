package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := numberOfSpecialSubstrings("pop")
	log.Printf("debug %v", output)
}
