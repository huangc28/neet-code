package main

import "log"

func mergeTriplets(triplets [][]int, target []int) bool {
	left, right := 0, 1
	for right < len(triplets) {
		leftTriplet := triplets[left]
		if !validTriplet(leftTriplet, target) {
			log.Printf("debug 1%v", leftTriplet)
			left += 1
			right += 1
			continue
		}

		rightTriplet := triplets[right]
		if !validTriplet(rightTriplet, target) {
			right += 1
			continue
		}

		// merge both triplet
		mergedTriplet := merge(leftTriplet, rightTriplet)
		if eq(mergedTriplet, target) {
			return true
		}

		triplets[right] = mergedTriplet
		left = right
		right += 1
	}

	return false
}

func merge(a, b []int) []int {
	merged := make([]int, 3)
	for i := 0; i < len(a); i++ {
		merged[i] = max(a[i], b[i])
	}
	return merged
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func eq(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func validTriplet(t, target []int) bool {
	for idx, num := range t {
		if num > target[idx] {
			return false
		}
	}
	return true
}
