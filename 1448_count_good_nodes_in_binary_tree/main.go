package main

/*
Time: O(n)
Space: O(h) height of the tree. maximum recursive call stack is the height of the tree
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type GoodNodeCount struct {
	Count int
}

func goodNodes(root *TreeNode) int {
	gnc := &GoodNodeCount{0}
	goodNodesHelper(root, root, gnc)
	return gnc.Count
}

func goodNodesHelper(root, maxValNode *TreeNode, gnc *GoodNodeCount) {
	if root == nil {
		return
	}

	if maxValNode.Val <= root.Val {
		gnc.Count += 1
	}

	if root.Val > maxValNode.Val {
		maxValNode = root
	}

	goodNodesHelper(root.Left, maxValNode, gnc)
	goodNodesHelper(root.Right, maxValNode, gnc)
}
