package binarytree

import "strconv"

var res []string
var paths []int

func binaryTreePaths(root *TreeNode) []string {
	res = make([]string, 0)
	paths = make([]int, 0)
	trace10(root)
	return res
}

func trace10(root *TreeNode) {
	if root == nil {
		return
	}
	paths = append(paths, root.Val)
	defer func() {
		paths = paths[:len(paths)-1]
	}()
	if root.Left == nil && root.Right == nil { // 子节点
		// 加入到res
		join()
		return
	}
	trace10(root.Left)
	trace10(root.Right)
}

func join() {
	var s string
	for i := 0; i < len(paths); i++ {
		s += strconv.Itoa(paths[i])
		if i != len(paths)-1 {
			s += "->"
		}
	}
	res = append(res, s)
}
