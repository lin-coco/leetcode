package main

func reverseWords(s string) string {
	// 双指针
	n := len(s)
	res := ""
	l := 0
	r := 0
	for l < n && r < n {
		for l < n && s[l] == ' ' {
			l++
		}
		if l < n {
			r = l
		} else {
			break
		}
		for r+1 < n && s[r+1] != ' ' {
			r++
		}
		if l <= r {
			res = " " + string(s[l:r+1]) + res
		}
		l = r + 1
	}
	return res[1:]
}
