package backtrack

func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	paths = make([]int, 0)
	targetSum = n
	sum = 0
	tracking(1, k)
	return res
}

var res [][]int
var paths []int
var targetSum int
var sum int

func tracking(start, k int) {
	if k == 0 {
		if sum == targetSum {
			res = append(res, copySlice(paths))
		}
		return
	}
	for i := start; i <= 9; i++ {
		paths = append(paths, i)
		sum += i
		if sum <= targetSum {
			tracking(i+1, k-1)
		}
		sum -= i
		paths = paths[:len(paths)-1]
	}
}

func copySlice(slice []int) []int {
	size := len(slice)
	t := make([]int, size)
	for i := 0; i < size; i++ {
		t[i] = slice[i]
	}
	return t
}
