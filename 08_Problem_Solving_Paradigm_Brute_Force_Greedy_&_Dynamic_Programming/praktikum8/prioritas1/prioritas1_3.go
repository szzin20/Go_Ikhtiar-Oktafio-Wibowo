package main

import "fmt"

var memo = make(map[int]int)

func fibonacci(number int) int {
	// Memeriksa apakah hasil perhitungan sudah ada di memo
	if val, ok := memo[number]; ok {
		return val
	}
	
	// Kasus dasar untuk angka Fibonacci 0 dan 1
	if number <= 1 {
		return number
	}

	// Menghitung angka Fibonacci ke-n dengan rekursi
	result := fibonacci(number-1) + fibonacci(number-2)

	// Menyimpan hasil perhitungan di memo
	memo[number] = result

	return result
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}