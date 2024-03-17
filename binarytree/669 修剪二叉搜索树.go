package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	return trace50(root, low, high)
}

func trace50(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		root = trace50(root.Right, low, high)
		return root
	} else if root.Val > high {
		root = trace50(root.Left, low, high)
		return root
	} else {
		root.Left = trace50(root.Left, low, high)
		root.Right = trace50(root.Right, low, high)
		return root
	}
}
