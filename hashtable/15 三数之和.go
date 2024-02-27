package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	// 排序
	// -4 -1 -1 0 1 2
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			if t := nums[i] + nums[left] + nums[right]; t < 0 {
				left++
			} else if t > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right {
					left++
					if nums[left] != nums[left-1] {
						break
					}
				}
			}
		}
	}
	return res
}
