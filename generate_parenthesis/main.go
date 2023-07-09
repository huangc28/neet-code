package main

import "strings"

/*
for n = 3

we need to have 3 open parenthesis and 3 close parenthesis

openCount   = 0
closedCount = 0

- we can only add closed parenthesis only when closedCount < openCount

   [(]
   / \
[()] [((]
*/

type CountInfo struct {
	OpenCount  int
	CloseCount int
}

func generateParenthesis(n int) []string {
	// Remember the solution for each recursive branch.
	stack := make([]string, 0)
	res := make([]string, 0)

	var backtrack func(int, int)
	backtrack = func(openN, closeN int) {
		// base case, found a solution
		if openN == n && closeN == n {
			sol := strings.Join(stack, "")
			res = append(res, sol)
		}

		// add open parenthesis
		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closeN)
			stack = stack[:len(stack)-1]
		}

		if closeN < openN {
			stack = append(stack, ")")
			backtrack(openN, closeN+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return res

	//countInfo := &CountInfo{0, 0}
	//sol := Helper(countInfo, n)
	//return sol
}

//func Helper(ci *CountInfo, numPairs int) []string {
//wellformed := []string{}

//// Add open parenthesis only when OpenCount < numPairs
//// and

////if ci.CloseCount

//return []string{}
//}
