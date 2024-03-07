package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	trace5(root, &res)
	return res
}
func trace5(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	trace5(root.Left, res)
	*res = append(*res, root.Val)
	trace5(root.Right, res)
}
