package binarytree

// var paths []int
var pathSum int

func hasPathSum(root *TreeNode, targetSum int) bool {
	paths = make([]int, 0)
	pathSum = 0
	return trace13(root, targetSum)
}

func trace13(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	paths = append(paths, root.Val)
	pathSum += root.Val
	defer func() {
		pathSum -= root.Val
		paths = paths[:len(paths)-1]
	}()
	if root.Left == nil && root.Right == nil { // 叶子结点
		if pathSum == targetSum {
			return true
		}
	}
	return trace13(root.Left, targetSum) || trace13(root.Right, targetSum)
}
