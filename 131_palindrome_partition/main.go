package main

/*
	aab

/   |  \
a   aa   aab
|    |
a    b
|
b
*/
func partition(s string) [][]string {
	ans := make([][]string, 0)
	part := make([]string, 0)

	var dfs func(i int, part []string)
	dfs = func(i int, part []string) {
		if i == len(s) {
			npart := make([]string, len(part))
			copy(npart, part)
			ans = append(ans, npart)
			return
		}

		// i is the base index where we will start
		// our substring from.
		for j := i; j < len(s); j++ {
			substr := s[i : j+1]
			if isPal(substr) {
				part = append(part, substr)
				dfs(j+1, part)
				part = part[:len(part)-1]
			}
		}
	}

	dfs(0, part)
	return ans
}

func isPal(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
