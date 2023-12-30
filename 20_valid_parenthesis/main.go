package main

/*
- use a stack,
- push open bracket the stack.
- when we encounter a matching close parentheses, pop that pop it out
- when we encounter a unmatched parentheses, return false
- if we still have left over parenthesis in the stack, the string is invalid.

example 1:

s = ()[]{}
         i
stack=""

example 2:

s = (]
     i
stack=(]

example
s = ]
    i
stack=

*/

var matches = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]

		// if char is an open parenthesis, push it to stac
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)

			continue
		}

		// if char is a close parenthesis, we check if it matches the top
		// open parenthesis in the stack, if not, return false, if yes, pop
		// the top parenthesis out from the stac
		expectedOpenParenthesis := matches[char]
		if len(stack) > 0 && expectedOpenParenthesis == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
