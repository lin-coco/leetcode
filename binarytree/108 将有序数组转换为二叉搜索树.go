package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	// 重新构建
	return trace51(nums)

}

func trace51(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := len(nums) / 2
	node := &TreeNode{Val: nums[midIndex]}
	node.Left = trace51(nums[:midIndex])
	node.Right = trace51(nums[midIndex+1:])
	return node
}
