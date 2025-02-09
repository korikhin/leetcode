package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetricBranches(root.Left, root.Right)
}

// Same function from "100. Same Tree"
func symmetricBranches(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val {
		return false
	}

	// ...but the leafs are swapped
	return symmetricBranches(a.Left, b.Right) && symmetricBranches(a.Right, b.Left)
}
