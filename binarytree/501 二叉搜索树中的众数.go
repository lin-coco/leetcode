package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	list1 = make(map[int]int)
	max = 0
	trace(root)
	res := make([]int, 0)
	for k, v := range list {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

var list1 map[int]int
var max int

func trace30(root *TreeNode) {
	if root == nil {
		return
	}
	trace(root.Left)
	list[root.Val] += 1
	if list[root.Val] > max {
		max = list[root.Val]
	}
	trace(root.Right)
}
