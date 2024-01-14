package main

import "strconv"

var operandsMap = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

func evalRPN(tokens []string) int {
	stack := make([]string, 0)

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if !operandsMap[token] {
			stack = append(stack, token)
			continue
		}

		operator := token
		arg1 := stack[len(stack)-2]
		arg2 := stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		res := operate(operator, arg1, arg2)
		stack = append(stack, res)
	}

	ans, _ := strconv.Atoi(stack[0])

	return ans
}

func operate(operator, arg1, arg2 string) string {
	arg1Int, _ := strconv.Atoi(arg1)
	arg2Int, _ := strconv.Atoi(arg2)

	res := 0

	switch operator {
	case "+":
		res = arg1Int + arg2Int
	case "-":
		res = arg1Int - arg2Int
	case "*":
		res = arg1Int * arg2Int
	case "/":
		res = arg1Int / arg2Int
	default:
		return ""
	}

	resStr := strconv.Itoa(res)

	return resStr
}
