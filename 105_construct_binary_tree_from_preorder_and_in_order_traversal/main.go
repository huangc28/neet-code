package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
preorder: root -> left -> right
inorder: left -> root -> right

	 3
	/ \

9   20

pre: [3]
in: [9]

in[0] 是 pre[0] 的 child
in[1] 是 pre[1] 的 child

正常情況下 currNode.Left = in[n]。如果 pre[n] == in[n-1]，代表 pre[n] 沒有 leaf node 了，我們要 traverse right

這邊可以知道 pre 會比 in 慢一個 step。當 pre[n] == in[n-1] 的話代表我們掉頭了。set currNode = in[n],  currNode.Right = pre[n+1],
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	midIdx := findMidIdx(inorder, root.Val)

	root.Left = buildTree(preorder[1:midIdx+1], inorder[:midIdx])
	root.Right = buildTree(preorder[midIdx+1:], inorder[midIdx+1:])

	return root
}

func findMidIdx(arr []int, v int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return i
		}
	}

	return -1
}

func buildTree_2(preorder []int, inorder []int) *TreeNode {
	// we use preorder's first element as the root element.
	root := &TreeNode{}
	root.Val = preorder[0]

	// we find the index of first element value in inorder arr.
	// split, left and right
	inLeft, inRight := findAndSplit(root.Val, inorder)

	// based on the length of left and right slice from in-order arr,
	// we slice left / right with the same length from preorder arr
	// to construct subtree nodes
	if len(inLeft) == 0 {
		root.Left = nil
	} else {
		root.Left = buildTree(preorder[1:1+len(inLeft)], inLeft)
	}

	if len(inRight) == 0 {
		root.Right = nil
	} else {
		root.Right = buildTree(preorder[1+len(inLeft):], inRight)
	}

	return root
}

func findAndSplit(v int, inorder []int) ([]int, []int) {
	idx := 0
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == v {
			break
		}
	}
	return inorder[:idx], inorder[idx+1:]
}
