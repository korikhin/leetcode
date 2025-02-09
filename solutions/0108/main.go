package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	n /= 2
	root := &TreeNode{Val: nums[n]}

	root.Left = sortedArrayToBST(nums[:n])
	root.Right = sortedArrayToBST(nums[n+1:])

	return root
}
