package main

import (
	"log"
	"testing"
)

type Node struct {
	Val int
}

func TestCase1(t *testing.T) {
	arr := append([]*Node{}, &Node{Val: 1})
	arr = append(arr, &Node{Val: 2})
	arr = arr[:1]
	log.Printf("debug arr %v", arr[1])
}
