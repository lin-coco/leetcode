package backtrack

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	paths = make([]int, 0)
	tracking(1, n, k)
	return res
}

var res [][]int
var paths []int

func tracking(start, end int, k int) {
	if k == 0 { // res
		res = append(res, copySlice(paths))
		return
	}
	for i := start; i <= end; i++ {
		paths = append(paths, i)
		tracking(i+1, end, k-1)
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
