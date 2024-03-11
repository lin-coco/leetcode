package binarytree

var list []int

func isValidBST(root *TreeNode) bool {
	list = make([]int, 0)
	trace20(root)
	if len(list) <= 1 {
		return true
	}
	for i := 1; i < len(list); i++ {
		if list[i-1] >= list[i] {
			return false
		}
	}
	return true
}
func trace20(root *TreeNode) {
	if root == nil {
		return
	}
	trace(root.Left)
	list = append(list, root.Val)
	trace(root.Right)
}
