package main

/*
We need to compare 3 things at each node:

1. left = max(left, 0)
2. right = max(right, 0)
3. maxChild = max(left, right)
4. maxVal = max(maxVal, max(left+right+root, root+maxChild))
5. return root.Val + maxChild
*/
func BinaryTreeMaximumPath_2(root *TreeNode) int {
	var maxVal int

	var postorder func(root *TreeNode) int
	postorder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMax := max(postorder(root.Left), 0) // If it's negative number, we ignore it.
		rightMax := max(postorder(root.Right), 0)

		maxChild := max(leftMax, rightMax)

		// compare root+maxChild with root+leftMax+rightMax
		localMax := max(root.Val+maxChild, root.Val+rightMax+leftMax)

		// compare localMax with global max
		maxVal = max(maxVal, localMax)

		return root.Val + maxChild
	}

	postorder(root)

	return maxVal
}
