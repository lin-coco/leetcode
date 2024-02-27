package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // 元素:索引
	for i := 0; i < len(nums); i++ {
		e := nums[i]
		if v, b := m[target-e]; b {
			return []int{v, i}
		}
		m[e] = i
	}
	return []int{}
}
