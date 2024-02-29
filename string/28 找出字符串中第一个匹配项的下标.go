package main

func strStr(haystack string, needle string) int {
	// kmp算法
	// 1. 生成模式串前缀表
	// 2. 循环文本串，回退前缀表
	n := len(needle)
	next := make([]int, n)
	getNext(next, needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func getNext(next []int, needle string) {
	j := 0
	next[0] = j
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[j] != needle[i] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
}
