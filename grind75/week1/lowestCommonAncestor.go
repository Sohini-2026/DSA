package week1

// LCA is where the search paths to p and q first diverge
func LowestcommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		return LowestcommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return LowestcommonAncestor(root.Right, p, q)
	}

	return root
}

func LowestcommonAncestor2(root, p, q *TreeNode) *TreeNode {
	curr := root

	for curr != nil {
		if curr.Val > p.Val && curr.Val > q.Val {
			curr = curr.Left
		} else if curr.Val < p.Val && curr.Val < q.Val {
			curr = curr.Right
		} else {
			return curr
		}
	}

	return nil
}
