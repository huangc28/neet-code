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
	if root == nil {
		return []int{}
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	ans := make([]int, 0)

	for len(q) > 0 {
		qLen := len(q)
		ans = append(ans, q[len(q)-1].Val)

		for i := 0; i < qLen; i++ {
			node := q[0]
			q = q[1:]

			// If node has left child, push to q
			if node.Left != nil {
				q = append(q, node.Left)
			}
			// If node has right child, push to q
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return ans
}
