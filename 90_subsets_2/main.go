package main

/*
nums = [1,2,2]
*/
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	subset := make([]int, 0)

	var backtrack func(i int, subset []int)
	backtrack = func(i int, subset []int) {
		if i == len(nums) {
			csubset := make([]int, len(subset))
			copy(csubset, subset)
			res = append(res, csubset)
			return
		}

		// consider nums[i]
		backtrack(i+1, append(subset, nums[i]))

		// don't consider nums[i]
		j := i
		for j < len(nums) && nums[j] == nums[i] {
			j += 1
		}
		backtrack(j, subset)
	}

	backtrack(0, subset)
	return res
}
