package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	isSameTree(root, subRoot)

	if root.Left != nil && isSameTree(root.Left, subRoot) {
		return true
	}

	if root.Right != nil && isSameTree(root.Right, subRoot) {
		return true
	}

	return false
}

func isSameTree(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil && subRoot != nil {
		return false
	} else if root != nil && subRoot == nil {
		return false
	}

	return root.Val == subRoot.Val && isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
}
