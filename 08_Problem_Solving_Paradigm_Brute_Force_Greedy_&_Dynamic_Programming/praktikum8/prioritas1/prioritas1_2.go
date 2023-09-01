package main

import (
	"fmt"
)

func segitigaPascal(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	segitiga := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		segitiga[i] = make([]int, i+1)
		segitiga[i][0], segitiga[i][i] = 1, 1
		for j := 1; j < i; j++ {
			segitiga[i][j] = segitiga[i-1][j-1] + segitiga[i-1][j]
		}
	}
	return segitiga
}

func main() {
	segitiga := segitigaPascal(5)
	fmt.Println(segitiga)
}