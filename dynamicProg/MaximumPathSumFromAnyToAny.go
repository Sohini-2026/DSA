package dynamicProg

import "github.com/Sohini-2026/DSA/baseDS"

func MaximumPathSum(root *baseDS.TreeNode) int {
	maxSum := -1 << 31 // Smallest int
	maxPathSumHelper(root, &maxSum)
	return maxSum
}

func maxPathSumHelper(root *baseDS.TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	left := max(0, maxPathSumHelper(root.Left, maxSum))   // Ignore negative paths
	right := max(0, maxPathSumHelper(root.Right, maxSum)) // Ignore negative paths

	// Update the global maximum if needed
	*maxSum = max(*maxSum, left+right+root.Value)

	// Return max path sum "starting" from this node
	return root.Value + max(left, right)
}

/*
      1
     / \
    2   3
   / \
  4   5

Step-by-Step Execution
1. Start at the root (1)
Call maxPathSumHelper(1, &maxSum)
2. Go to left child (2)
Call maxPathSumHelper(2, &maxSum)
3. Go to left child (4)
Call maxPathSumHelper(4, &maxSum)
Both left and right are nil, so both recursive calls return 0
left = 0, right = 0
Update: *maxSum = max(*maxSum, 0+0+4) = 4
Return: 4 + max(0, 0) = 4
4. Go to right child (5)
Call maxPathSumHelper(5, &maxSum)
Both left and right are nil, so both recursive calls return 0
left = 0, right = 0
Update: *maxSum = max(4, 0+0+5) = 5
Return: 5 + max(0, 0) = 5
5. Back at node (2)
left = 4 (from node 4)
right = 5 (from node 5)
Update: *maxSum = max(5, 4+5+2) = max(5, 11) = 11
Return: 2 + max(4, 5) = 2 + 5 = 7
6. Go to right child (3)
Call maxPathSumHelper(3, &maxSum)
Both left and right are nil, so both recursive calls return 0
left = 0, right = 0
Update: *maxSum = max(11, 0+0+3) = 11
Return: 3 + max(0, 0) = 3
  7. Back at root (1)
left = 7 (from node 2)
right = 3 (from node 3)
Update: *maxSum = max(11, 7+3+1) = max(11, 11) = 11
Return: 1 + max(7, 3) = 8
Final Result
The function returns maxSum = 11
The maximum path sum is the path: 4 → 2 → 5 (sum = 11)
*/
