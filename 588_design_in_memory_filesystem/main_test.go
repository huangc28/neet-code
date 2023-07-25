package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	fs := Constructor()
	fs.Ls("/")
	fs.Mkdir("/a/b/c")
}
