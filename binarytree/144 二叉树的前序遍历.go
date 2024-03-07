package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	trace3(root, &res)
	return res
}

func trace3(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	trace3(root.Left, res)
	trace3(root.Right, res)
}
