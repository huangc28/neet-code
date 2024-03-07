package main

func goodNodes_2(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode, maxVal int)
	dfs = func(root *TreeNode, maxVal int) {
		if root == nil {
			return
		}

		if root.Val >= maxVal {
			res += 1
		}

		dfs(root.Left, max(maxVal, root.Val))
		dfs(root.Right, max(maxVal, root.Val))
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
