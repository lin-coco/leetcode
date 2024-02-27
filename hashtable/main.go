package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{2, 5, 1, 3, 4}
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)
}
