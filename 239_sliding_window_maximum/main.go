package main

/*
nums = [1,3,-1,-3,5,3,6,7], k = 3

	i

q = 1-, 3-, -1-, -3-, 5-, 3-, 6-, 7
res = 3, 3, 5, 5, 6, 7
*/
func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0)
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// Pop right most element in the queue that is less than me
		for len(q) > 0 && q[len(q)-1] <= nums[i] {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if i+1 >= k {
			res = append(res, nums[q[0]])
		}

		// left element of queue might contain element
		// that is out of range that we don't need to consider
		if i-q[0]+1 >= k {
			q = q[1:]
		}
	}

	return res
}
