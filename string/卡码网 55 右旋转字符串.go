package main

import (
	"bufio"
	"os"
	"strconv"
)

func kama55() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	kStr, _ := reader.ReadString('\n')
	kStr = kStr[:len(kStr)-1]
	k, _ := strconv.Atoi(kStr)
	s, _ := reader.ReadString('\n')
	// 	s = s[:len(s)-1]
	n := len(s)
	k = k % n
	i := 0
	res := make([]byte, n)
	for i = 0; i < n-k; i++ {
		res[i+k] = s[i]
	}
	for i = n - k; i < n; i++ {
		res[i-n+k] = s[i]
	}
	writer.Write(res)
	writer.Flush()
}
