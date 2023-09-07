package main

import "testing"

func TestCase1(t *testing.T) {
	merge_2([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
