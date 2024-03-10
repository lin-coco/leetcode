package binarytree

func minDepth(root *TreeNode) int {
	return trace7(root)
}
func trace7(root *TreeNode) int {
	if root == nil {
		return 0
	}
	i := 1
	if root.Left == nil && root.Right == nil {
		return i
	} else if root.Left != nil && root.Right == nil {
		return i + trace7(root.Left)
	} else if root.Right != nil && root.Left == nil {
		return i + trace7(root.Right)
	} else {
		j := trace7(root.Left) // 2
		k := trace7(root.Right)
		if j < k {
			return i + j
		} else {
			return i + k
		}
	}
	return i
}
