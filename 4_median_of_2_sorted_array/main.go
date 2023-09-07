package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	shorterArray := nums1
	longerArray := nums2

	if len(shorterArray) > len(longerArray) {
		shorterArray, longerArray = longerArray, shorterArray
	}

	total := len(nums1) + len(nums2)
	half := (total + 1) / 2

	left := 0
	right := len(shorterArray) - 1

	for {

		// this works
		// mid := left + right
		// if left + right >= 0 {
		//     mid = (left + right) / 2
		// }

		// this also works
		mid := (left + right) >> 1

		shorterArrIdx := mid
		longerArrIdx := half - (mid + 1) - 1

		var shorterLeft, shorterRight float64
		var longerLeft, longerRight float64

		if shorterArrIdx >= 0 {
			shorterLeft = float64(shorterArray[shorterArrIdx])
		} else {
			shorterLeft = math.Inf(-1)
		}

		if shorterArrIdx+1 < len(shorterArray) {
			shorterRight = float64(shorterArray[shorterArrIdx+1])
		} else {
			shorterRight = math.Inf(1)
		}

		if longerArrIdx >= 0 {
			longerLeft = float64(longerArray[longerArrIdx])
		} else {
			longerLeft = math.Inf(-1)
		}

		if longerArrIdx+1 < len(longerArray) {
			longerRight = float64(longerArray[longerArrIdx+1])
		} else {
			longerRight = math.Inf(1)
		}

		if shorterLeft <= longerRight && longerLeft <= shorterRight {
			// found valid partition
			if total%2 == 0 { // even
				return (max(shorterLeft, longerLeft) + min(shorterRight, longerRight)) / 2
			}

			// odd
			return max(shorterLeft, longerLeft)
		} else if shorterLeft > longerRight {
			// shorter array has to find a mid that is less than longer right
			right = mid - 1
		} else {
			// shorter array has to find a mid that is larger than longer right
			left = mid + 1
		}
	}
}

func min(i, j float64) float64 {
	if i < j {
		return i
	}
	return j
}

func max(i, j float64) float64 {
	if i > j {
		return i
	}
	return j
}
