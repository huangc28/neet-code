package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	//output := searchMatrix(
	//[][]int{
	//{1, 3, 5, 7},
	//{10, 11, 16, 20},
	//{23, 30, 34, 60},
	//},
	//3,
	//)

	output := searchMatrix(
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		13,
	)

	log.Printf("debug 1 %v", output)
}
