package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	k %= n
	if k == 0 {
		return head
	}

	// Make it a Circle List
	tail.Next = head

	newTail := head
	for i := 1; i < n-k; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}
