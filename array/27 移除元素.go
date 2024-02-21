package main

/*
https://leetcode.cn/problems/remove-element/description/
此题使用左右指针
*/
func removeElement(nums []int, val int) int {
	// 如果cur元素 = val，则将last元素覆盖到cur元素
	size := len(nums)
	if size == 0 {
		return 0
	}
	cur := 0
	last := size - 1
	for cur <= last {
		for last >= 0 && nums[last] == val {
			last--
			size--
			nums = nums[:size]
		}
		if size <= 0 {
			return 0
		}
		if cur >= last {
			break
		}
		// nums[last] != val
		if nums[cur] == val {
			nums[cur] = nums[last]
			last--
			size--
			nums = nums[:size]
		}
		cur++
	}
	return size
}
