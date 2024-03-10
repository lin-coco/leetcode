package binarytree

func buildTree(inorder []int, postorder []int) *TreeNode {
	return trace14(inorder, postorder)
}

func trace14(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = trace14(inorder[:i], postorder[:i])
	root.Right = trace14(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}
