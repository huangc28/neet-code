package main

func rightSideView_2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	res := make([]int, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		res = append(res, queue[len(queue)-1].Val)
		newQueue := make([]*TreeNode, 0)

		// pop everything out from the queue
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}

			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}

		queue = append(queue, newQueue...)
	}

	return res
}
