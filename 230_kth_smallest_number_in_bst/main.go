package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type KthNum struct {
	K   int
	Num int
}

func kthSmallest(root *TreeNode, k int) int {
	kth := &KthNum{K: k, Num: 0}
	kthSmallestHelper(root, kth)
	return kth.Num
}

func kthSmallestHelper(root *TreeNode, kth *KthNum) {
	if root == nil {
		return
	}

	kthSmallestHelper(root.Left, kth)

	if kth.K > 0 {
		kth.K -= 1
		if kth.K == 0 {
			kth.Num = root.Val
		}
	}

	if kth.K == 0 {
		return
	}

	kthSmallestHelper(root.Right, kth)
}
