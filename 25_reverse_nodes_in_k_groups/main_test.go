package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	head := reverseKGroup(n1, 2)
	for head != nil {
		log.Printf("debug 1 %v", head.Val)
		head = head.Next
	}
	//log.Printf("debug 2 %v", res.Next)
	//log.Printf("debug 3 %v", res.Next.Next)
}
