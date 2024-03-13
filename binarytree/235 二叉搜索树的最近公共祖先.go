package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor40(root, p, q *TreeNode) *TreeNode {
	return trace40(root, p, q)
}

func trace40(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	t1 := trace40(root.Left, p, q)
	t2 := trace40(root.Right, p, q)
	if t1 != nil && t2 != nil {
		return root
	} else if t1 == nil && t2 != nil {
		return t2
	} else if t1 != nil && t2 == nil {
		return t1
	} else {
		return nil
	}
}
