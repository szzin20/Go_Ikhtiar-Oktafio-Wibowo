package main

import (
	"fmt"
	"strconv"
)

func generateBinary(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	return ans
}

func main() {
	n := 2
	ans := generateBinary(n)
	fmt.Println(n, "=", ans )

	n = 3
	ans = generateBinary(n)
	fmt.Println(n,"=", ans)
}
