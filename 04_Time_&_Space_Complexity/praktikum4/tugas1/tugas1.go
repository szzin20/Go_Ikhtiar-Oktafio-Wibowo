package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	
	if number <= 3 {
		return true
	}
	
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	
	// Menggunakan teorema bahwa semua bilangan prima memiliki bentuk 6k ± 1, di mana k adalah bilangan bulat.
	// hanya perlu memeriksa pembagi hingga akar kuadrat dari bilangan.
	maxDivisor := int(math.Sqrt(float64(number)))
	
	for i := 5; i <= maxDivisor; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}