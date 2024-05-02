package main

import "testing"

func TestCase1(t *testing.T) {
	n1 := &Node{Val: 7}
	n2 := &Node{Val: 13}
	n3 := &Node{Val: 11}
	n4 := &Node{Val: 10}
	n5 := &Node{Val: 1}

	n1.Next = n2
	n1.Random = nil

	n2.Next = n3
	n2.Random = n1

	n3.Next = n4
	n3.Random = n5

	n4.Next = n5
	n4.Random = n3

	n5.Next = nil
	n5.Random = n1

	head := n1

	copyRandomList2(head)
}
