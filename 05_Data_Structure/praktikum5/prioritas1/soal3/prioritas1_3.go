package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	countMap := make(map[int]int)
	var result []int

	for _, char := range angka {
		digit := int(char - '0')
		countMap[digit]++
	}

	for digit, count := range countMap {
		if count == 1 {
			result = append(result, digit)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}