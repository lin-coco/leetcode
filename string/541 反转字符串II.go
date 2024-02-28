package main

func reverseStr(s string, k int) string {
	// 模拟 双指针
	res := make([]byte, len(s))
	start := 0
	num := 0 // 计数
	for i := 0; i < len(s); i++ {
		num++
		if num == 2*k {
			l := start
			r := start + k - 1
			for l < r {
				res[l] = s[l] ^ s[r]
				res[r] = res[l] ^ s[r]
				res[l] = res[l] ^ res[r]
				l++
				r--
			}
			if l == r {
				res[l] = s[l]
			}
			for j := start + k; j < start+k+k; j++ {
				res[j] = s[j]
			}
			start += num
			num = 0
		}
	}
	if num <= k {
		l := start
		r := len(s) - 1
		for l < r {
			res[l] = s[l] ^ s[r]
			res[r] = res[l] ^ s[r]
			res[l] = res[l] ^ res[r]
			l++
			r--
		}
		if l == r {
			res[l] = s[l]
		}
	} else if num < 2*k {
		l := start
		r := start + k - 1
		for l < r {
			res[l] = s[l] ^ s[r]
			res[r] = res[l] ^ s[r]
			res[l] = res[l] ^ res[r]
			l++
			r--
		}
		if l == r {
			res[l] = s[l]
		}
		for i := start + k; i < len(s); i++ {
			res[i] = s[i]
		}
	}
	return string(res)
}
