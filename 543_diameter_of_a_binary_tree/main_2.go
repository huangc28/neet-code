package main

/*
每一個階層的 root 我們要判斷下面幾個 diameter

- 通過 root 的 diameter, (1+ left) + (1 + right). 為什麼要 +1? 因為是 edge
- 左邊 child node 的 diameter 跟右邊 child node 的 diameter. 哪個比較長？比較長的拿去跟 root diameter 比
- 上面兩個 diameter 哪個比較長？比較長的那個就是當前 root 的 max diameter
- 把當前 root diameter 跟 當前 max diameter 作比較。哪個比較長？update 當前 max diameter。

return 時，回傳 larger diameter between left and right diameter. so that parent node get the largest side diameter.
*/

func diameterOfBinaryTree_2(root *TreeNode) int {
	md := MaxDiameter{0}
	diameterOfBinaryTree_2helper(root, &md)
	return md.Diameter
}

func diameterOfBinaryTree_2helper(root *TreeNode, md *MaxDiameter) int {
	if root == nil {
		return -1
	}

	leftDiameter := 1 + diameterOfBinaryTree_2helper(root.Left, md)
	rightDiameter := 1 + diameterOfBinaryTree_2helper(root.Right, md)
	longerChildDiameter := max(leftDiameter, rightDiameter)
	md.Diameter = max(md.Diameter, max(leftDiameter+rightDiameter, longerChildDiameter))

	return longerChildDiameter
}
