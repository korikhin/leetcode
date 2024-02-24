package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	carry := dummy

	for head != nil && head.Next != nil {
		// Swapped pair
		first, second := head, head.Next

		// Set second element first
		carry.Next = second

		// Prevents infinite loop by linking first node to rest of list after swap
		first.Next = second.Next

		// Set first element
		second.Next = first

		// Move to the next pair
		carry, head = first, first.Next
	}

	return dummy.Next
}
