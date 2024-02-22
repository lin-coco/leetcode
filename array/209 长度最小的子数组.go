package main

func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口
	// 双指针ij指向0和1，小了j++，大了i++
	size := len(nums)
	if size == 1 {
		if nums[0] < target {
			return 0
		} else {
			return 1
		}
	}
	i := 0
	j := 0
	sum := nums[0]
	res := -1
	for {
		if j < size-1 {
			if sum < target {
				j++
				sum += nums[j]
			} else {
				t := j - i + 1
				if t < res || res == -1 {
					res = t
				}
				sum -= nums[i]
				i++
			}
		} else {
			if sum < target {
				break
			} else {
				t := j - i + 1
				if t < res || res == -1 {
					res = t
				}
				sum -= nums[i]
				i++
			}
		}
	}
	if res == -1 {
		res = 0
	}
	return res
}
