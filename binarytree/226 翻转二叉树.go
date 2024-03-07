package binarytree

func invertTree(root *TreeNode) *TreeNode {
	trace(root)
	return root
}
func trace(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	trace(root.Left)
	trace(root.Right)
}
