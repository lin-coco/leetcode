package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 后续遍历
	target, _ := trace31(root, p, q)
	return target
}

func trace31(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	n1, b1 := trace31(root.Left, p, q)
	n2, b2 := trace31(root.Right, p, q)
	if !b1 && !b2 { // 都为false
		if root.Val == p.Val || root.Val == q.Val {
			return nil, true
		} else {
			return nil, false
		}
	} else if b1 && b2 { // 都为true
		return root, true
	} else { // 只有一个为true
		if n1 != nil {
			return n1, true
		}
		if n2 != nil {
			return n2, true
		}
		// 两个都为nil,其中一个为true
		if root.Val == p.Val || root.Val == q.Val {
			return root, true
		} else {
			return nil, true
		}
	}
	// if root.Val == p.Val || root.Val == q.Val {

	// }
	// if b1 && b2 {
	//     return root,true
	// } else if (b1 || b2) && (root.Val == p.Val || root.Val == q.Val) {
	//     return root,true
	// } else {
	//     return nil,false
	// }
}
