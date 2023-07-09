package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	reorderList(head)
	log.Printf("debug ans %v", head)
}

func TestReorderList2Case1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	reorderList2(head)
	for head != nil {
		log.Printf("debug ans %v", head.Val)
		head = head.Next
	}
}
