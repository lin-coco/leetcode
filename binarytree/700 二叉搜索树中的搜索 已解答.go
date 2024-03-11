package binarytree

func searchBST(root *TreeNode, val int) *TreeNode {
	return trace23(root, val)
}

func trace23(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if val < root.Val {
		return trace23(root.Left, val)
	} else {
		return trace23(root.Right, val)
	}
}
