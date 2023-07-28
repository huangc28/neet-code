package main

import (
	"strings"
)

type Codec struct{}

func (c *Codec) Encode(strs []string) string {
	builder := strings.Builder{}
	for _, str := range strs {
		builder.WriteByte(byte(len(str)))
		builder.WriteString(str)
	}
	return builder.String()
}

func (c *Codec) Decode(strs string) []string {
	res := make([]string, 0)
	lenP := 0
	for lenP < len(strs) {
		length := int(strs[lenP])
		res = append(res, strs[lenP+1:lenP+1+length])
		lenP = lenP + length + 1
	}
	return res
}
