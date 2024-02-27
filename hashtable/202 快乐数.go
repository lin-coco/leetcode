package main

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	m := make(map[int]struct{})
	m[n] = struct{}{}
	for {
		sum := 0 // 平方和
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if sum == 1 {
			return true
		}
		if _, e := m[sum]; e {
			return false
		}
		m[sum] = struct{}{}
		n = sum
	}
}
