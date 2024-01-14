package main

/*
 solution 1:
temperatures=[73,74,75,71,69,72,76,73]
              i

iterate from i ~ n, find the first index kth that has warmer temperature than ith day.
if `temperature[kth] > temperature[ith]`, add kth-ith to ans. break the iteration

This is a O(n^2) time complexity solution


solution 2:

Have a stack that push the ith day. Iterate through the temperatures.
If temperatures[ith] day > temperatures[stack[top]], we find a warmer day.
set ans[stack[top]] = ith-stack[top], and pop top of the stack. push ith in stack
If temperatures[ith] day <= temperatures[stack[top]], we push ith in stack and contiunue iteration to the next day

temps=[73,74,75,71,69,72,76,73]
                          i
        0  1  2  3  4  5  6  7

stack: [~0, ~1, 2, ~3, ~4, ]
ans: [1, 1, 4, 2, 1, 0, 0, 0]
*/

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))

	stack := make([]int, 0, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 || temperatures[i] <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return nil
}
