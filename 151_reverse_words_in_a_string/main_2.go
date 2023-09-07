package main

/*
2 pointers solution
*/
func reverseWords_2(s string) string {
	s = reverseStr(s, 0, len(s)-1)
	l, r := 0, 0

	for r < len(s) && s[r] == ' ' {
		r += 1
	}

	for r < len(s) {
		// right pointer search for end of current word.
		for r+1 < len(s) && s[r+1] != ' ' {
			r += 1
		}

		// now we locate start and end pos of current string
		// reverse it.
		s = reverseStr(s, l, r)

		// move left forward. This will be the correct place
		// where we want to place next word.
		for l < len(s) && s[l] != ' ' {
			l += 1
		}
		l += 1

		// move right forward until the start of the next
		// word.
		r = l
		for r < len(s) && s[r] == ' ' {
			r += 1
		}
	}

	// trim ending spaces
	ep := len(s) - 1
	for ; ep >= 0; ep-- {
		if s[ep] != ' ' {
			break
		}
	}

	return s[:ep+1]
}

func reverseStr(s string, start, end int) string {
	l, r := start, end
	runes := []rune(s)
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}
