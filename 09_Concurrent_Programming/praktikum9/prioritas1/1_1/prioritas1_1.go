package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Println(i)
		}
		time.Sleep(3 * time.Second) // Tunggu selama 3 detik
	}
}

func main() {
	x := 5 
	go printMultiples(x)

	// Biarkan goroutine berjalan selamanya
	select {}
}