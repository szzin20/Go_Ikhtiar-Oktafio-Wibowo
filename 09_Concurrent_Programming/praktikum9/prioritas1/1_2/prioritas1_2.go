package main

import (
	"fmt"
	"time"
)

func printMultiples3(ch chan<- int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			ch <- i // Mengirim bilangan kelipatan 3 ke channel
		}
		time.Sleep(3 * time.Second) // Tunggu selama 3 detik
	}
}

func main() {
	ch := make(chan int)
	go printMultiples3(ch)

	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		}
	}
}