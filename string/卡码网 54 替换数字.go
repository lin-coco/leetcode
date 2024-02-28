package main

import (
	"bufio"
	"os"
)

func kama54() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	res := make([]byte, 0, len(input))
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			res = append(res, 'n', 'u', 'm', 'b', 'e', 'r')
		} else {
			res = append(res, input[i])
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.Write(res)
	writer.Flush()
}
