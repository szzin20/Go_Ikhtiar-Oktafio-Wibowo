package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	diagonal1Sum := 0
	diagonal2Sum := 0

	for i := 0; i < len(matrix); i++ {
		diagonal1Sum += matrix[i][i]
		diagonal2Sum += matrix[i][len(matrix)-1-i]
	}

	absoluteDifference := int(math.Abs(float64(diagonal1Sum - diagonal2Sum)))
	fmt.Println("jumlah diagonal 1 : ", diagonal1Sum)
	fmt.Println("jumlah diagonal 2 : ", diagonal2Sum)
	fmt.Printf("Selisih absolut antara jumlah diagonal: %d\n", absoluteDifference)
}