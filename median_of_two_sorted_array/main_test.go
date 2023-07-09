package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{}
	median1 := findMedianSortedArrays_2(nums1, nums2)
	log.Printf("debug ans %v", median1)
}

func TestShiftNegativeNumber(t *testing.T) {
	log.Printf("debug %v", -1>>1)
}
