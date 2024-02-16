package main

import "testing"

func TestListNode_String(t *testing.T) {
	tests := []struct {
		name     string
		list     *ListNode
		expected string
	}{
		{
			name:     "empty list",
			list:     nil,
			expected: "[]",
		},
		{
			name:     "single element",
			list:     &ListNode{Val: 1},
			expected: "[1]",
		},
		{
			name:     "multiple elements",
			list:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			expected: "[1, 2, 3]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.String(); got != tt.expected {
				t.Errorf("ListNode.String() = %s, want %s", got, tt.expected)
			}
		})
	}
}

func newList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func listsEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	return l1 == nil && l2 == nil
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "both empty",
			l1:       []int{},
			l2:       []int{},
			expected: []int{},
		},
		{
			name:     "first empty",
			l1:       []int{},
			l2:       []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "second empty",
			l1:       []int{1, 2, 3},
			l2:       []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "no carry",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "with carry",
			l1:       []int{9, 9, 9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := newList(tt.l1)
			l2 := newList(tt.l2)
			result := AddTwoNumbers(l1, l2)
			expected := newList(tt.expected)
			if !listsEqual(result, expected) {
				t.Errorf("Test '%s' failed: expected %s, got %s", tt.name, expected, result)
			}
		})
	}
}
