package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "{}"
	valid := isValid(s)
	log.Printf("debug valid %v", valid)
}
