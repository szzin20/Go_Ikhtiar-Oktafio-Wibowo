package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, mu *sync.Mutex) int {
	defer wg.Done()

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	// Mengunci akses ke variabel result sebelum memperbarui nilainya
	mu.Lock()
	defer mu.Unlock()
	return result
}

func main() {
	n := 5 
	var wg sync.WaitGroup
	var mu sync.Mutex
	var results = make([]int, n+1)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(i int) {
			results[i] = factorial(i, &wg, &mu)
		}(i)
	}

	wg.Wait()

	// Mencetak hasil faktorial
	for i, result := range results {
		fmt.Printf("%d! = %d\n", i, result)
	}
}