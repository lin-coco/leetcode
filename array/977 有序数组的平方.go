package main

/*
两种思路，都用双指针（左右指针）
1. 指针指向左右两边，向内遍历，依次放入res数组（从最右边放入）
2. 先找到绝对值最小的那一个索引x，再从x-1和x+1双指针相反方向遍历比较，依次放入res数组（从最左边放入）
主要的切入点是指针的位置
*/
func sortedSquares(nums []int) []int {
	// 先找到绝对值最小的那一个，再从x-1和x+1双指针相反方向遍历比较，依次放入res数组
	size := len(nums)
	res := make([]int, size)
	if nums[0] >= 0 {
		// 都是递增
		for i := 0; i < size; i++ {
			res[i] = nums[i] * nums[i]
		}
		return res
	}
	// 有负数
	min := 0
	for i := 1; i < size; i++ {
		t := nums[i]
		if nums[i] < 0 {
			t = -t
		}
		if t < -nums[min] {
			min = i
		}
		if nums[i] > 0 {
			break
		}
	}
	res[0] = nums[min] * nums[min]
	left := min - 1
	right := min + 1
	i := 1
	for left >= 0 || right < size {
		if left < 0 {
			res[i] = nums[right] * nums[right]
			i++
			right++
		} else if right >= size {
			res[i] = nums[left] * nums[left]
			i++
			left--
		} else {
			if (-nums[left]) < nums[right] {
				res[i] = nums[left] * nums[left]
				i++
				left--
			} else {
				res[i] = nums[right] * nums[right]
				i++
				right++
			}
		}
	}
	return res
}
