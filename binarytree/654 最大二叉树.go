package binarytree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return trace21(nums)
}

func trace21(nums []int) *TreeNode {
	if size := len(nums); size == 0 {
		return nil
	} else if size == 1 {
		return &TreeNode{Val: nums[0]}
	}
	// 找出最大值
	mIndex := maxIndex(nums)
	root := &TreeNode{Val: nums[mIndex]}
	root.Left = trace21(nums[:mIndex])
	root.Right = trace21(nums[mIndex+1:])
	return root
}

func maxIndex(nums []int) int {
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}
	return max
}
