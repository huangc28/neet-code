package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var kth int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		if k > 0 {
			k--
			if k == 0 {
				kth = root.Val
				return
			}
		}

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	return kth
}
