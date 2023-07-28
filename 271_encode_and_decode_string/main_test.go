package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	codec := Codec{}
	enc1 := codec.Encode([]string{"hello", "world"})
	dec1 := codec.Decode(enc1)
	log.Printf("debug 2 %v", dec1)
}
