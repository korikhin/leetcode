package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return head
	}

	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return head.Next
		}
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
