package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if m[nums1[i]] == 1 {
			continue
		}
		m[nums1[i]]++
	}
	res := make(map[int]struct{}, 0)
	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] == 1 {
			res[nums2[i]] = struct{}{}
		}
	}
	l := make([]int, 0, len(res))
	for k := range res {
		l = append(l, k)
	}
	return l
}
