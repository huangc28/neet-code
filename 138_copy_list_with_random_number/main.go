package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
- Create copied list from original list with no random pointer and create a copied node map.
- Iterate original and copied at the same time, let the value of the original node tells you what the random pointer should be for the copied node
*/
func copyRandomList(head *Node) *Node {
	orig := head
	var (
		copied     *Node
		copiedHead *Node
	)
	copiedMap := make(map[*Node]*Node)

	for orig != nil {
		copiedNode := &Node{Val: orig.Val}
		copiedMap[orig] = copiedNode
		if copied == nil {
			copied = copiedNode
			copiedHead = copied
		} else {
			copied.Next = copiedNode
			copied = copied.Next
		}
		orig = orig.Next
	}

	curr1 := head
	curr2 := copiedHead

	for curr1 != nil {
		if curr1.Random == nil {
			curr2.Random = nil
		} else {
			curr2.Random = copiedMap[curr1.Random]
		}

		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return copiedHead
}
