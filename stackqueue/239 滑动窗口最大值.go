package stackqueue

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	// 栈
	stack := make([]int, 0, k)
	for i := 0; i < k; i++ {
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	res = append(res, stack[0])
	for i := k; i < len(nums); i++ {
		if nums[i-k] == stack[0] {
			stack = stack[1:]
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
		res = append(res, stack[0])
	}
	return res
}
