package week1

/** * Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := InvertBinaryTree(root.Left)
	right := InvertBinaryTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}

func InvertBinaryTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr.Left, curr.Right = curr.Right, curr.Left

		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	return root
}
func InvertBinaryTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		curr.Left, curr.Right = curr.Right, curr.Left

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}

		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return root
}

func PrintTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		result = append(result, curr.Val)

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}

		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return result
}
