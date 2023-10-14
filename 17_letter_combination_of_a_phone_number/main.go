package main

var digMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	comb := make([]byte, 0)

	var dfs func(i int, comb []byte)
	dfs = func(i int, comb []byte) {
		if i == len(digits) {
			substr := string(comb)
			ans = append(ans, substr)
			return
		}

		dig := digits[i]
		letters := digMap[dig]

		for j := 0; j < len(letters); j++ {
			comb = append(comb, letters[j])
			dfs(i+1, comb)
			comb = comb[:len(comb)-1]
		}
	}

	dfs(0, comb)

	return ans
}
