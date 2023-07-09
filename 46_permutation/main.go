package main

func removeAtIndex(s []int, idx int) []int {
	ns := make([]int, len(s)-1, cap(s))
	copy(ns, s[:idx])
	copy(ns[idx:], s[idx+1:])
	return ns

}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	var cb func(comb, candidates []int)
	cb = func(comb, candidates []int) {
		if len(comb) == len(nums) {
			ncomb := make([]int, len(comb))
			copy(ncomb, comb)
			res = append(res, ncomb)
			return
		}

		for i := 0; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			cb(comb, removeAtIndex(candidates, i))
			comb = comb[:len(comb)-1]
		}
	}

	cb(comb, nums)

	return res
}

func permute_2(nums []int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	visited := make(map[int]int)

	var cb func(comb []int, visited map[int]int)
	cb = func(comb []int, visited map[int]int) {
		if len(comb) == len(nums) {
			ncomb := make([]int, len(comb))
			copy(ncomb, comb)
			res = append(res, ncomb)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] == 0 {
				visited[i] += 1
				comb = append(comb, nums[i])
				cb(comb, visited)
				comb = comb[:len(comb)-1]
				visited[i] -= 1
			}
		}
	}

	cb(comb, visited)
	return res
}
