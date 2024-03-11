package binarytree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return trace22(root1, root2)
}

func trace22(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	root := &TreeNode{}
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil {
		root.Val = root2.Val
		root.Left = trace22(nil, root2.Left)
		root.Right = trace22(nil, root2.Right)
		return root
	} else if root2 == nil {
		root.Val = root1.Val
		root.Left = trace22(root1.Left, nil)
		root.Right = trace22(root1.Right, nil)
		return root
	}
	root.Val = root1.Val + root2.Val
	root.Left = trace22(root1.Left, root2.Left)
	root.Right = trace22(root1.Right, root2.Right)
	return root
}
