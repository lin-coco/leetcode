package main

func generateMatrix(n int) [][]int {
	// 模拟
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}
	var i, j int // 二维数组索引
	t := 1       // 放入的元素值
	u := n - 1   // 每行(列)填入的数量
	for u > 0 {
		for k := 0; k < u; k++ {
			nums[i][j] = t
			t++
			j++
		}
		for k := 0; k < u; k++ {
			nums[i][j] = t
			t++
			i++
		}
		for k := 0; k < u; k++ {
			nums[i][j] = t
			t++
			j--
		}
		for k := 0; k < u; k++ {
			nums[i][j] = t
			t++
			i--
		}
		u -= 2
		i++
		j++
	}
	if u == 0 {
		nums[i][j] = t
	}
	return nums
}
