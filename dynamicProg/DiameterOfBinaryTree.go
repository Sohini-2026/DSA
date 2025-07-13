package dynamicProg

import "github.com/Sohini-2026/DSA/baseDS"

func DiameterOfBinaryTree(root *baseDS.TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	leftHeight := DiameterOfBinaryTree(root.Left, res)
	rightHeight := DiameterOfBinaryTree(root.Right, res)

	// Update the diameter if the path through this node is larger
	*res = max(*res, leftHeight+rightHeight+1)

	// Return the height of the current node
	return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
