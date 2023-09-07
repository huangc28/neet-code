package main

func rotate(nums []int, k int) {
	// locate the index of pivot where the rotation starts
	stepsToRotate := k % len(nums)
	pivot := len(nums) - stepsToRotate
	newArr := append(nums[pivot:], nums[:pivot]...)

	for i := 0; i < len(nums); i++ {
		nums[i] = newArr[i]
	}
}

func rotate_2(nums []int, k int) {
	stepsToMove := k % len(nums)

	reverse(nums, 0, len(nums)-1)

	// reverse first k elements
	reverse(nums, 0, stepsToMove-1)

	// reverse the rest of elements
	reverse(nums, stepsToMove, len(nums)-1)
}

func rotate_3(nums []int, k int) {
	temp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp[(i+k)%len(nums)] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = temp[i]
	}
}

func reverse(nums []int, start, end int) {
	l, r := start, end
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
