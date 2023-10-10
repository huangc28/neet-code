package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	//output := alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"})
	//log.Printf("debug %v", output)

	output2 := alienOrder([]string{"z", "z"})
	log.Printf("debug 2 %v", output2)
}
