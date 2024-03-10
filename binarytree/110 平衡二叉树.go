package binarytree

func isBalanced(root *TreeNode) bool {
	_, b := trace9(root)
	return b
}

// 计算深度
func trace9(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	i := 1
	j, b1 := trace9(root.Left)
	k, b2 := trace9(root.Right)
	c := 0
	if j < k {
		c = i + k
	} else {
		c = i + j
	}
	if j-k < -1 || j-k > 1 || !b1 || !b2 {
		return c, false
	} else {
		return c, true
	}
}
