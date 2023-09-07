package main

/*
Time: m(length of word list) * n(length of string)
space: O(n) height of recursive tree which is wordDict
*/
func wordBreak_topdown(s string, wordDict []string) bool {
	dp := make(map[int]bool)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if v, ok := dp[i]; ok {
			return v
		}

		if i >= len(s) {
			return true
		}

		for _, word := range wordDict {
			wl := len(word)
			if i+wl <= len(s) && s[i:i+wl] == word { // word in dict is found in s
				exists := dfs(i + wl)
				// cache the result of finding word in dict from s[i:wl]
				// it is possible that i+wl can not be found in word dict.
				// we cache that in dp[i+wl]=false. The next time we visit
				// dp[i+wl] we can just return the result from the cache instead
				dp[i+wl] = exists
				if exists {
					return true
				}
			}
		}
		return false
	}
	return dfs(0)
}

/*
s = leetcode // len: 8
wordDict=["leet", "code"]
dp = [         ]

dp[8] = true // empty string matches every string in the word dict
dp[7] = false // "e" not in dict
dp[6] = false // "de" not in dict
dp[5] = false // "ode" not in dict
dp[4] = true // "code" in dict, whenever we found a match, reset word
dp[3] = false // "t" not in dict
dp[2] = false // "et" not in dict
dp[1] = false // "eet" not in dict
dp[0] = dp[len(word)]
*/
func wordBreak_bottom_up(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)

	dp[len(s)] = true

	// previous matched position
	mp := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 {
			for _, word := range wordDict {
				if s[0:mp] == word {
					dp[0] = dp[len(word)]
				}
			}
		} else {
			for _, word := range wordDict {
				if s[i:mp] == word {
					dp[i] = true
					mp = i
				}
			}
		}
	}

	return dp[0]
}
