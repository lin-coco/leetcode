package main

import "fmt"

func main() {
	nums := make([]int, 5)
	nums = nums[:0]
	fmt.Println(len(nums))
}
