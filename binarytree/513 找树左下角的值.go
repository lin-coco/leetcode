package binarytree

func findBottomLeftValue(root *TreeNode) int {
	// 层序遍历
	// list := make([]*TreeNode,0)
	// list = append(list,root)
	// size := 1
	// for i := 0i < size;i++ {
	//     if list[0].Left
	// }
	dep = 1
	val = root.Val
	trace12(root, 1)
	return val
}

var dep int
var val int

// 递归遍历
func trace12(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	trace12(root.Left, depth+1)
	if depth > dep {
		dep = depth
		val = root.Val
	}
	trace12(root.Right, depth+1)
}
