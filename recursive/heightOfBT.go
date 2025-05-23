package recursive

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func FindHeightOfBT(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := FindHeightOfBT(root.Left)
	rightHeight := FindHeightOfBT(root.Right)

	return 1 + max(leftHeight, rightHeight)
}
