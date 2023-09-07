package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val >= root.Val && q.Val <= root.Val {
		return root
	}

	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}

	if p.Val <= root.Val && q.Val <= root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val >= root.Val && q.Val >= root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return nil
}

func lowestCommonAncestor_2(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor_2(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor_2(root.Left, p, q)
	}

	return root
}
