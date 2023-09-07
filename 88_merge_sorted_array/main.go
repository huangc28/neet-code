package main

import (
	"log"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m++
	}
	sort.Ints(nums1)
}

func merge_2(nums1 []int, m int, nums2 []int, n int) {
	np1, np2 := 0, 0
	for np1 < m && np2 < n {
		if nums1[np1] <= nums2[np2] {
			np1++
		} else {
			nums1[np1], nums2[np2] = nums2[np2], nums1[np1]
		}
		np1++
		np2++
	}

	log.Printf("debug 1 %v", nums1)
	log.Printf("debug 2 %v", nums2)

	// merge nums1 and num2
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m++
	}
}
