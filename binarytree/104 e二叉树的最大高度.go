package binarytree

func maxDepth(root *TreeNode) int {
	return trace6(root)
}

func trace6(root *TreeNode) int {
	if root == nil {
		return 0
	}
	i := 1
	j := trace6(root.Left)
	k := trace6(root.Right)
	if j > k {
		return i + j
	} else {
		return i + k
	}
}
