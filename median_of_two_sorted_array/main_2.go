package main

import (
	"log"
	"math"
)

/*
longerArray:  [1, 2, 3, 4, 5]

	p

shorterArray: [3, 4]

	m

totalLength = 7
half = 3

先找到 shorterArray 的 mid 我們才能找到 longer array 要取多少個 element 當 longerSubArray。

shorterSubArray + longerSubarray 組成 leftPartition。

我們怎麼知道這個 partition 是正確的 partition?

shorterSubArray := shorterArray[:m+1]
longerSubArray := longerArray[:half-len(shorterSubArray)]

longerSubArray [1, 2]
shorterSubArray [3]

if shorterSubArray[len(shorterSubArray)-1] <= longerArray[half-len(shorterSubArray)] &&

	longerSubArray[len(longerSubArray)-1] <= shorterArray[m+1] {
		// valid partition
	}
*/
func findMedianSortedArrays_2(nums1 []int, nums2 []int) float64 {
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
		mid := left + right
		if mid >= 0 {
			mid = (left + right) / 2
		}

		// or
		// mid := (left +right) >> 1

		shorterArrIdx := mid
		longerArrIdx := half - (mid + 1) - 1

		var shorterLeft, shorterRight float64
		var longerLeft, longerRight float64

		log.Printf("debut 1 %v %v", shorterArray, shorterArrIdx)

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

//func findMedianSortedArrays_2(nums1 []int, nums2 []int) float64 {
//A, B := nums1, nums2
//total := len(nums1) + len(nums2)
//half := (total + 1) / 2

//var Aleft, Aright float64
//var Bleft, Bright float64

//if len(B) < len(A) {
//A, B = B, A
//}

//l, r := 0, len(A)-1
//for {
//i := (l + r) >> 1 // A
//j := half - i - 2 // B

//log.Printf("debug 1 %v %v", l, r)
//log.Printf("debug 1 %v", i)

//if i >= 0 {
//Aleft = float64(A[i])
//} else {
//Aleft = math.Inf(-1)
//}

//if (i + 1) < len(A) {
//Aright = float64(A[i+1])
//} else {
//Aright = math.Inf(1)
//}

//if j >= 0 {
//Bleft = float64(B[j])
//} else {
//Bleft = math.Inf(-1)
//}

//if (j + 1) < len(B) {
//Bright = float64(B[j+1])
//} else {
//Bright = math.Inf(1)
//}

//// partition is correct
//if Aleft <= Bright && Bleft <= Aright {
//// odd
//if total%2 == 1 {
//return max(Aleft, Bleft)
//}
//// even
//return (max(Aleft, Bleft) + min(Aright, Bright)) / 2
//} else if Aleft > Bright {
//r = i - 1
//} else {
//l = i + 1
//}
//}
//}
