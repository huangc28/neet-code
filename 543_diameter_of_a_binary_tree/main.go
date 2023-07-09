package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MaxDiameter struct {
	Diameter int
}

func diameterOfBinaryTree(root *TreeNode) int {
	md := &MaxDiameter{Diameter: 0}
	diameterOfBinaryTreeHelper(root, md)
	return md.Diameter
}

func diameterOfBinaryTreeHelper(root *TreeNode, md *MaxDiameter) int {
	// we are at empty leaf node, it's length is 0 since no edges
	// are connected to them
	if root == nil {
		return -1
	}

	left := 1 + diameterOfBinaryTreeHelper(root.Left, md)
	right := 1 + diameterOfBinaryTreeHelper(root.Right, md)
	longerChild := max(left, right)

	// compare max d with longer child
	currNodeMaxD := max(longerChild, left+right)

	// compare max d with left + right
	md.Diameter = max(md.Diameter, currNodeMaxD)

	return longerChild
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
