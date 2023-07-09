package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	lower := math.MinInt
	upper := math.MaxInt
	return isValidBSTHelper(root, lower, upper)
}

func isValidBSTHelper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	isLeftValid := isValidBSTHelper(root.Left, lower, root.Val)
	isRightValid := isValidBSTHelper(root.Right, root.Val, upper)

	return isLeftValid && isRightValid
}
