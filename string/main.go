package main

import "fmt"

func main() {
	//fmt.Println(reverseStr("abcdefg", 1))
	kama55()
}

func test() {
	var input string
	fmt.Scan(&input)
	res := make([]byte, 0, len(input))
	for i := 0; i < len(input); i++ {
		if res[i] >= '0' && res[i] <= '9' {
			res = append(res, 'n', 'u', 'm', 'b', 'e', 'r')
		} else {
			res = append(res, input[i])
		}
	}
	fmt.Println(string(res))
}
