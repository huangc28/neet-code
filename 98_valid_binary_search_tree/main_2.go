package main

import "math"

func isValidBST_2(root *TreeNode) bool {
	maxVal := math.MaxInt
	minVal := math.MinInt

	var dfs func(root *TreeNode, maxVal, minVal int) bool
	dfs = func(root *TreeNode, maxVal, minVal int) bool {
		if root == nil {
			return true
		}

		if root.Val >= minVal || root.Val <= maxVal {
			return false
		}

		return dfs(root.Left, root.Val, minVal) || dfs(root.Right, maxVal, root.Val)
	}

	return dfs(root, maxVal, minVal)
}
