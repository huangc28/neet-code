package main

import "fmt"

/*
example: (*)
f(index, open)

	 f(0,0)
	   |
	f(1,1)
	 /    |      \

f(2,0)  f(2,2) f(2,1)

	          |
	        f(3,1)      \
			 f(3,0)
*/
func checkValidString(s string) bool {
	dp := make(map[string]bool)
	var dfs func(i int, openNum int) bool
	dfs = func(i, openNum int) bool {
		if v, ok := dp[fmt.Sprintf("%d_%d", i, openNum)]; ok {
			return v
		}

		// base case if we are at the end
		// of the string, and if the number of
		// open parenthesis is 0, we found a valid
		// string
		if i == len(s) {
			return openNum == 0
		}

		valid := false

		// encounter '('
		if s[i] == '(' {
			valid = dfs(i+1, openNum+1)
		} else if s[i] == ')' && openNum > 0 {
			// encounter ')', if we still have enough
			// open parenthesis to deduct, keep looking
			// into recursion
			valid = dfs(i+1, openNum-1)

		} else {
			// encounter '*', we have 3 variations '(' or ')' or just '*'
			valid = dfs(i+1, openNum+1)
			if !valid && openNum > 0 {
				valid = dfs(i+1, openNum-1)
			}
			if !valid {
				valid = dfs(i+1, openNum)
			}

		}

		dp[fmt.Sprintf("%d_%d", i, openNum)] = valid
		return dp[fmt.Sprintf("%d_%d", i, openNum)]
	}

	return dfs(0, 0)
}
