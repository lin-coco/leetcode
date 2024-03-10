package binarytree

func sumOfLeftLeaves(root *TreeNode) int {
	return trace11(root, false)
}

func trace11(root *TreeNode, left bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil && left {
		return root.Val
	}
	i := trace11(root.Left, true)
	j := trace11(root.Right, false)
	return i + j
}
