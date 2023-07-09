package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Determine which array is smaller
	smallerArr := nums1
	greaterArr := nums2

	var smallerLeft, smallerRight float64
	var greaterLeft, greaterRight float64

	if len(nums1) > len(nums2) {
		smallerArr = nums2
		greaterArr = nums1
	}

	totalLength := len(smallerArr) + len(greaterArr)
	mid := totalLength / 2

	pLeft := 0
	pRight := len(smallerArr) - 1

	for {
		smallerArrMidIndex := (pRight - pLeft) / 2
		// greaterArrMidIndex := (mid - (smallerArrMidIndex + 1)) - 1
		greaterArrMidIndex := mid - smallerArrMidIndex - 2

		if smallerArrMidIndex >= 0 {
			smallerLeft = float64(smallerArr[smallerArrMidIndex])
		} else {
			smallerLeft = math.Inf(-1)
		}

		if smallerArrMidIndex+1 < len(smallerArr) {
			smallerRight = float64(smallerArr[smallerArrMidIndex+1])
		} else {
			smallerRight = math.Inf(1)
		}

		if greaterArrMidIndex >= 0 {
			greaterLeft = float64(greaterArr[greaterArrMidIndex])
		} else {
			greaterLeft = math.Inf(-1)
		}

		if greaterArrMidIndex+1 < len(greaterArr) {
			greaterRight = float64(greaterArr[greaterArrMidIndex+1])
		} else {
			greaterRight = math.Inf(1)
		}

		if smallerLeft <= greaterRight && greaterLeft <= smallerRight {
			if totalLength%2 == 1 {
				return max(smallerLeft, greaterLeft)
			}

			return (max(smallerLeft, greaterLeft) + min(smallerRight, greaterRight)) / 2
		} else if smallerLeft > greaterRight {
			pRight = smallerArrMidIndex - 1
		} else {
			pLeft = smallerArrMidIndex + 1
		}
	}
}

func max(i, j float64) float64 {
	if i > j {
		return i
	}
	return j
}

func min(i, j float64) float64 {
	if i < j {
		return i
	}
	return j
}
