package main

import "sort"

/*
nums = [2,5,2,1,2] target = 5
ordered = [1,2,2,2,5]
res: [1,2,2]

i = 0

[1] --> 0 ~ 4

	[1,2] --> 1 ~ 4
	  [1,2,2] ok --> 2 ~ 4
	  [1,2,5]
	[1,5]

[2] --> 1 ~ 4

	[2,2] --> 2 ~ 4
	 [2,2,2] --> 3 ~ 4
	 [2,2,5] --> 3 ~ 4
	[]

這題也是用 decision tree backtracking 來解
*/

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	comb := make([]int, 0)
	sort.Ints(candidates)

	var cb func(i int, sum int, comb []int)
	cb = func(i, sum int, comb []int) {
		if sum == target {
			ncomb := make([]int, len(comb))
			copy(ncomb, comb)
			ans = append(ans, ncomb)
			return
		}

		if i == len(candidates) {
			return
		}

		// pick candidate at i
		if candidates[i]+sum <= target {
			comb = append(comb, candidates[i])
			cb(i+1, candidates[i]+sum, comb)
			comb = comb[:len(comb)-1]
		}

		// skip all numbers that equal to candidates[i]
		j := i
		for j < len(candidates) && candidates[j] == candidates[i] {
			j++
		}

		cb(j, sum, comb)
	}

	cb(0, 0, comb)

	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	res := make([][]int, 0)
	comb := make([]int, 0)

	var backtracking func(idx int, sum int, comb []int)
	backtracking = func(idx int, sum int, comb []int) {
		if sum == target {
			ncomb := make([]int, len(comb))
			copy(ncomb, comb)
			res = append(res, ncomb)
			return
		} else if idx == len(candidates) || sum > target {
			return
		}

		// take candidates[idx] into consideration
		backtracking(idx+1, sum+candidates[idx], append(comb, candidates[idx]))

		// don't take candidates[idx] into consideration
		j := idx
		for j < len(candidates) && candidates[j] == candidates[idx] {
			j += 1
		}
		backtracking(j, sum, comb)
	}

	backtracking(0, 0, comb)

	return res
}
