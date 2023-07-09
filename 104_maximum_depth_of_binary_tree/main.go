package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := 1 + maxDepth(root.Left)
	rightDepth := 1 + maxDepth(root.Right)

	return max(leftDepth, rightDepth)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
