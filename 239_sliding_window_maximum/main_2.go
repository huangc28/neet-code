package main

/*
	 nums = [1,3,-1,-3,5,3,6,7], k = 3
	           l
		      r

deq = [~0,1,2]
sol = [3]
*/
func maxSlidingWindow_2(nums []int, k int) []int {
	deq := make([]int, 0, len(nums))
	res := make([]int, 0, len(nums))

	l := 0
	for r := 0; r < len(nums); r++ {
		for len(deq) > 0 && nums[r] > nums[deq[len(deq)-1]] {
			// pop the right most number in nums
			deq = deq[:len(deq)-1]
		}
		deq = append(deq, r)

		if (r - l + 1) < k {
			continue
		}

		res = append(res, nums[deq[0]])

		// move l forward until (r-l+1) == k-1
		l += 1
		if l > deq[0] {
			deq = deq[1:]
		}

		// more secure
		// for (r - l + 1) > k-1 {
		// 	l += 1 // move l pointer forward
		// 	if len(deq) > 0 && l > deq[0] {
		// 		deq = deq[1:]
		// 	}
		// }
	}

	return res
}
