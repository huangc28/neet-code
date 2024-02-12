package main

import "math"

func findMedianSortedArray2(nums1 []int, nums2 []int) float64 {
	shorter, longer := nums1, nums2
	if len(shorter) > len(longer) {
		shorter, longer = longer, shorter
	}
	// Get the total number of element in nums1 and nums2
	total := len(nums1) + len(nums2)
	// 14, half = 7
	// 15, half = 7, but we want to include the middle position
	// thus, we want half to have 8 element
	half := (total + 1) / 2

	left, right := 0, len(shorter)-1

	for {
		shorterIdx := (left + right) >> 2
		// longerIdx so we know the right partition to split longer slice
		longerIdx := half - (shorterIdx + 1) - 1

		var shorterRight, shorterLeft, longerRight, longerLeft int
		if shorterIdx >= 0 {
			shorterRight = shorter[shorterIdx]
		} else {
			shorterRight = -math.MaxInt
		}

		if shorterIdx+1 < len(shorter) {
			shorterLeft = shorter[shorterIdx+1]
		} else {
			shorterLeft = math.MaxInt
		}

		if longerIdx >= 0 {
			longerRight = longer[longerIdx]
		} else {
			longerRight = -math.MaxInt
		}

		if longerIdx+1 < len(longer) {
			longerLeft = longer[longerIdx+1]
		} else {
			longerLeft = math.MaxInt
		}

		if shorterRight <= longerLeft && longerRight <= shorterLeft {
			// we found the correct partition
			if (len(shorter)+len(longer))%2 == 0 {
				// even number of elements
				return float64(max2(longerRight, shorterRight)+min2(longerLeft, shorterLeft)) / 2
			}

			// odd number of elements
			return float64(max2(longerRight, shorterRight))
		}

		if longerRight > shorterLeft {
			// expand shorter, shrink longer by searching more left in shorter slice
			left = shorterIdx + 1
		} else {
			// expand longer, shrink shorter by searching more right in shorter slice
			right = shorterIdx - 1
		}
	}
}

func max2(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min2(i, j int) int {
	if i < j {
		return i
	}
	return j
}
