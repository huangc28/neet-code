package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MaxSum struct {
	Val int
}

func maxPathSum(root *TreeNode) int {
	maxSum := &MaxSum{Val: root.Val}
	maxPathSumHelper(root, maxSum)
	return maxSum.Val
}

func maxPathSumHelper(root *TreeNode, maxSum *MaxSum) int {
	if root == nil {
		return 0
	}

	leftMaxSum := maxPathSumHelper(root, maxSum)
	rightMaxSum := maxPathSumHelper(root, maxSum)

	// we don't need to consider negative number
	leftMaxSum = max(leftMaxSum, 0)
	rightMaxSum = max(rightMaxSum, 0)

	maxSum.Val = max(maxSum.Val, root.Val+leftMaxSum+rightMaxSum)

	return root.Val + max(leftMaxSum, rightMaxSum)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
