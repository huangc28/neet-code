package main

/*
11106
i
 i
  i
*/

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	var dfs func(i int) int
	dfs = func(i int) int {
		if dp[i] != 0 {
			return dp[i]
		} else if s[i] == '0' {
			return 0
		}

		// current char is valid, move a step forward
		num := dfs(i + 1)
		if (i+1 < len(s)) && s[i] == '1' || s[i] == '2' && (s[i+1] >= 0 && s[i+1] <= 6) {
			num += dfs(i + 2)
		}
		dp[i] = num
		return dp[i]
	}

	return dfs(0)
}
