package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	trace4(root, &res)
	return res
}

func trace4(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	trace4(root.Left, res)
	trace4(root.Right, res)
	*res = append(*res, root.Val)
}
