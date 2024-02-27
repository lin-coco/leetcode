package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 两两for循环
	n := len(nums1)
	res := 0
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[nums1[i]+nums2[j]]++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v, b := m[0-nums3[i]-nums4[j]]; b {
				res += v
			}
		}
	}
	return res
}
