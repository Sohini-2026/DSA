package dynamicProg

import "github.com/Sohini-2026/DSA/baseDS"

func MaximumPathSumLeafToLeaf(root *baseDS.TreeNode) int {
	maxSum := -1 << 31 // Smallest int (acts as negative infinity)
	maxPathSumHelperLeafToLeaf(root, &maxSum)
	return maxSum
}

func maxPathSumHelperLeafToLeaf(root *baseDS.TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Value
	}

	left := maxPathSumHelperLeafToLeaf(root.Left, maxSum)
	right := maxPathSumHelperLeafToLeaf(root.Right, maxSum)

	// Only update maxSum if both children exist (leaf-to-leaf path)
	if root.Left != nil && root.Right != nil {
		*maxSum = max(*maxSum, left+right+root.Value)
		return root.Value + max(left, right)
	}

	// If only one child exists, return path sum including that child
	if root.Left == nil {
		return root.Value + right
	}

	return root.Value + left
}
