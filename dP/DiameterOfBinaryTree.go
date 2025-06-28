package dp

import "github.com/Sohini-2026/DSA/baseDS"

func DiameterOfBinaryTree(root *baseDS.TreeNode, res *int) {
	if root == nil {
		return
	}

	leftHeight := 0
	rightHeight := 0

	// Calculate height of left subtree
	DiameterOfBinaryTree(root.Left, res)
	if root.Left != nil {
		leftHeight = root.Left.Value
	}

	// Calculate height of right subtree
	DiameterOfBinaryTree(root.Right, res)
	if root.Right != nil {
		rightHeight = root.Right.Value
	}

	// Update the diameter if needed
	*res = max(*res, leftHeight+rightHeight)

	// Return the height of the current node
	root.Value = max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
