package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	// 中序递归
	list2 = make([]int, 0)
	trace(root)
	min := list[1] - list[0]
	for i := 2; i < len(list); i++ {
		if v := list[i] - list[i-1]; v < min {
			min = v
		}
	}
	return min
}

var list2 []int

func trace32(root *TreeNode) {
	if root == nil {
		return
	}
	trace(root.Left)
	list = append(list, root.Val)
	trace(root.Right)

}
