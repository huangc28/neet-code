package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec2 struct{}

func (c *Codec2) Encode(strs []string) string {
	encstr := ""

	for i := 0; i < len(strs); i++ {
		str := strs[i]
		encstr += fmt.Sprintf("%d:%s", len(str), str)
	}

	return encstr
}

func (c *Codec2) Decode(strs string) []string {
	start := 0
	res := make([]string, 0)
	for start < len(strs) {
		// try to find position of `:`
		columnIdx := start + strings.Index(strs[start:], ":")

		// find the length of decoded string
		strLen, _ := strconv.Atoi(strs[start:columnIdx])

		// slice the decoded string from `strs`
		decstr := strs[columnIdx+1 : columnIdx+1+strLen]

		res = append(res, decstr)

		// update position of start for next decode
		start = columnIdx + 1 + strLen
	}

	return res
}
