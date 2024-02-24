package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var v []string
	for n != nil {
		v = append(v, fmt.Sprint(n.Val))
		n = n.Next
	}
	return fmt.Sprintf("[%s]", strings.Join(v, ", "))
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}

	for head := dummy; l1 != nil || l2 != nil || carry > 0; head = head.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		head.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}

	return dummy.Next
}
