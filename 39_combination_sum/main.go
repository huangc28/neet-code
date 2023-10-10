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

func combinationSum_2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	comb := make([]int, 0)

	var dfs func(idx int, target int, comb []int)
	dfs = func(idx int, target int, comb []int) {
		if target == 0 {
			ncomb := make([]int, len(comb))
			copy(ncomb, comb)
			ans = append(ans, ncomb)
			return
		}

		if idx == len(candidates) || target < 0 {
			return
		}

		// pick the current index
		if candidates[idx] <= target {
			comb = append(comb, candidates[idx])
			dfs(idx, target-candidates[idx], comb)
			comb = comb[:len(comb)-1]
		}

		// don't pick the current index, we don't consider current index on next recursion!
		dfs(idx+1, target, comb)

	}

	dfs(0, target, comb)

	return ans
}

/*
[2,3,6,7] target=7


       f(0,[2])
 f(0,[2,2])
 f(0,[2,2,2]) f(1,[2,2,3]) f(2,[2,2,6]) f(2,[2,2,7])


*/

func combinationSum_3(candidates []int, target int) [][]int {
	comb := make([]int, 0)
	ans := make([][]int, 0)

	var dfs func(i int, sum int, comb []int)
	dfs = func(i int, sum int, comb []int) {
		if sum == target {
			ncomb := make([]int, len(comb))
			copy(ncomb, comb)
			ans = append(ans, ncomb)
			return
		}

		if i >= len(candidates) {
			return
		}

		if sum+candidates[i] <= target {
			comb = append(comb, candidates[i])
			dfs(i, sum+candidates[i], comb)
			comb = comb[:len(comb)-1]
		}

		dfs(i+1, sum, comb)
	}

	dfs(0, 0, comb)
	return ans
}
