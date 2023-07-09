package main

import (
	"strconv"
)

/*
["2","1","+","3","*"]

Iterate through tokens. When we encounter operator pop first 2 operands out from the stack.
Perform operation of on the 2 operands, push the result back to stack.
*/

var operatorMaps = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if _, isOperator := operatorMaps[token]; isOperator {
			opt1 := stack[len(stack)-1]
			opt2 := stack[len(stack)-2]

			stack = stack[:len(stack)-2]
			stack = append(stack, operate(token, opt2, opt1))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}

func operate(operator string, opt2, opt1 int) int {
	switch operator {
	case "+":
		return opt2 + opt1
	case "-":
		return opt2 - opt1
	case "*":
		return opt2 * opt1
	case "/":
		return opt2 / opt1
	default:
		return 0
	}
}
