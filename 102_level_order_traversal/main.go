package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	sol := make([][]int, 0)

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		levelSol := make([]int, 0)
		newQueue := make([]*TreeNode, 0)

		// empty the current queue
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			levelSol = append(levelSol, node.Val)

			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}

			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}

		queue = append(queue, newQueue...)
		sol = append(sol, levelSol)
	}

	return sol
}
