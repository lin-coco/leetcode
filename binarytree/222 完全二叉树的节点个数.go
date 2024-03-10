package binarytree

func countNodes(root *TreeNode) int {
	return trace8(root)
}

func trace8(root *TreeNode) int {
	if root == nil {
		return 0
	}
	i := 1
	return i + trace8(root.Left) + trace8(root.Right)
}
