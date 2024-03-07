package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 层序遍历
	res := make([][]int, 0)
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	flag := true
	size := len(nodes)
	for flag {
		resElem := make([]int, 0)
		flag = false
		s := 0
		for i := 0; i < size; i++ {
			n := nodes[i]
			resElem = append(resElem, n.Val)
			if n.Left != nil {
				flag = true
				nodes = append(nodes, n.Left)
				s += 1
			}
			if n.Right != nil {
				flag = true
				nodes = append(nodes, n.Right)
				s += 1
			}
		}
		res = append(res, resElem)
		nodes = nodes[size:]
		size = s
	}
	return res
}
