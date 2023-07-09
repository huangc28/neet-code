package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
我們從左邊的 subtree 依序加入到 queue 中。queue 的最後一個 node 就是站在右邊的人可以看到的

res:
1, 3

queue:

~1, ~2, ~3, ~5, 4
*/
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		canViewNode := queue[len(queue)-1]
		res = append(res, canViewNode.Val)

		// append subtree to queue from left to right
		newLevelNodes := make([]*TreeNode, 0)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				newLevelNodes = append(newLevelNodes, node.Left)
			}

			if node.Right != nil {
				newLevelNodes = append(newLevelNodes, node.Right)
			}
		}

		queue = newLevelNodes
	}

	return res
}
