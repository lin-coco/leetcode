package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	var parent *TreeNode
	var oper int // 1left 2right
	for {
		if cur == nil {
			return root
		}
		if key < cur.Val {
			parent = cur
			oper = 1
			cur = cur.Left
		} else if key > cur.Val {
			parent = cur
			oper = 2
			cur = cur.Right
		} else {
			r, b := del(cur, parent, oper)
			if b {
				return r
			} else {
				return root
			}
		}
	}
	return root
}

func del(oldRoot *TreeNode, parent *TreeNode, oper int) (*TreeNode, bool) {
	var root *TreeNode
	node := oldRoot.Right
	if node == nil {
		root = oldRoot.Left
	} else {
		oldNode := node
		var prev *TreeNode
		for node.Left != nil {
			prev = node
			node = node.Left
		}
		if prev != nil {
			prev.Left = node.Right
		}
		root = node
		root.Left = oldRoot.Left
		if prev != nil {
			root.Right = oldNode
		}
	}
	oldRoot.Left = nil
	oldRoot.Right = nil
	if oper == 1 {
		parent.Left = root
		return nil, false
	} else if oper == 2 {
		parent.Right = root
		return nil, false
	} else {
		return root, true
	}
}
