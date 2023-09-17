package main

func jump_2(nums []int) int {
	l, r, res := 0, 0, 0
	for r < len(nums)-1 {
		farthest := 0
		for i := l; i <= r; i++ {
			farthest = max(farthest, i+nums[i])
		}

		l = r + 1
		r = farthest
		res++
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
