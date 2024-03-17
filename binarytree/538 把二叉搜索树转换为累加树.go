package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	// 先得出总和
	sum := traceSum(root)
	//再中序遍历
	trace52(root, &sum)
	return root
}

func traceSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	i := traceSum(root.Left)
	j := traceSum(root.Right)
	return i + j + root.Val
}

func trace52(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	trace52(root.Left, sum)
	oldVal := root.Val
	root.Val = *sum
	*sum -= oldVal
	trace52(root.Right, sum)
}
