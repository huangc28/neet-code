package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	output := numDecodings("11106")
	log.Printf("debug %v", output)
}
