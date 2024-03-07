package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return trace2(root.Left, root.Right)
}

func trace2(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right != nil {
		return false
	} else if left.Val != right.Val {
		return false
	} else {
		// 对称的 // 比较内侧 // 比较外侧
		return trace2(left.Right, right.Left) && trace2(left.Left, right.Right)
	}
}
