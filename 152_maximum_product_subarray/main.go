package main

import "math"

/*
[2,3,-2,4]
 r      l

largest: 4,6
leftProd : 1,2,6  ,-12,-48
rightProd: 1,4,-8 ,-24,-48

[-2,0,-1]
       l
  r

largest: min,-1,0
leftProd : 1,-2,0,-1
rightProd: 1,-1,0,-2

*/

func maxProduct(nums []int) int {
	leftProd, rightProd := 1, 1
	largest := math.MinInt32

	for i := 0; i < len(nums); i++ {
		if leftProd == 0 {
			leftProd = 1
		}

		if rightProd == 0 {
			rightProd = 1
		}

		leftProd *= nums[i]
		rightProd *= nums[len(nums)-i-1]

		largest = max(largest, max(leftProd, rightProd))
	}

	return largest
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
