package main

func reverseString(s []byte) {
	// 双指针交换
	l := 0
	r := len(s) - 1
	for l < r {
		s[l] = s[l] ^ s[r]
		s[r] = s[l] ^ s[r]
		s[l] = s[l] ^ s[r]
		l++
		r--
	}
}
