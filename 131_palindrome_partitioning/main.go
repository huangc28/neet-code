package main

func partition(s string) [][]string {
	res := make([][]string, 0)
	comb := make([]string, 0)

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx >= len(s) {
			ncomb := make([]string, len(comb))
			copy(ncomb, comb)
			res = append(res, ncomb)
			return
		}

		for i := idx; i < len(s); i++ {
			// check if s[:i+1] is palindrome or not
			if isPalindrome(s[idx : i+1]) {
				// push s[:i+1] to potential solution
				comb = append(comb, s[idx:i+1])
				dfs(i + 1)
				comb = comb[:len(comb)-1]
			}
		}
	}

	dfs(0)

	return res
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
