package main

type TreeNode struct {
	Right *TreeNode
	Left  *TreeNode
	Val   int
}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := isBalancedHelper(root)
	return isBalanced
}

// Post order traversal
func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	// check if left child is balanced and retrieve it's height
	leftHeight, leftBalanced := isBalancedHelper(root.Left)

	if !leftBalanced {
		return 0, false
	}

	// check if left child is balanced and retrieve it's height
	rightHeight, rightBalanced := isBalancedHelper(root.Right)

	if !rightBalanced {
		return 0, false
	}

	if absDiff(rightHeight, leftHeight) > 1 {
		return 0, false
	}

	return max(rightHeight, leftHeight) + 1, true
}

func absDiff(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
