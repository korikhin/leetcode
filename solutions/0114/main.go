package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var (
		tail *TreeNode
		dfs  func(node *TreeNode)
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		dfs(node.Left)

		node.Left = nil
		node.Right, tail = tail, node
	}

	dfs(root)
}
