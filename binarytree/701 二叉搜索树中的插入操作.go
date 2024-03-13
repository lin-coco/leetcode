package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	target := &TreeNode{Val: val}
	if root == nil {
		return target
	}
	var prev *TreeNode
	var oper int // 1: left 2: right
	current := root
	for {
		if current == nil {
			if oper == 1 {
				prev.Left = target
				return root
			} else {
				prev.Right = target
				return root
			}
		}
		prev = current
		if val < current.Val {
			oper = 1
			current = current.Left
		} else if val > current.Val {
			oper = 2
			current = current.Right
		}
	}
}
