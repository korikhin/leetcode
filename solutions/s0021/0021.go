package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			carry.Next = list1
			list1 = list1.Next
		} else {
			carry.Next = list2
			list2 = list2.Next
		}
		carry = carry.Next
	}

	if list1 != nil {
		carry.Next = list2
	} else if list2 == nil {
		carry.Next = list1
	}

	return dummy.Next
}
