package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	var cb func(index, sum int, comb []int)
	cb = func(index, sum int, comb []int) {
		if sum == target {
			newComb := make([]int, len(comb))
			copy(newComb, comb)
			res = append(res, newComb)
			return
		} else if sum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			cb(i, sum+candidates[i], comb)
			comb = comb[:len(comb)-1]
		}
	}

	cb(0, 0, comb)

	return res
}
