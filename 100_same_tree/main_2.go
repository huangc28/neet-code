package main

func isSameTree_2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) && (p != nil && q == nil) {
		return false
	}

	if !isSameTree_2(p.Left, q.Left) && !isSameTree_2(p.Right, q.Right) {
		return false
	}

	return p.Val == q.Val
}
