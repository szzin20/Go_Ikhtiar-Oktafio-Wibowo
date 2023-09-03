package main

import (
	"fmt"
	"sync"
)

func countLetters(text string, ch chan<- map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()

	counts := make(map[rune]int)
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			counts[char]++
		}
	}

	ch <- counts
}

func main() {
	text := "Hello, World! This is a sample text with letter frequency counting."

	// Membagi teks menjadi beberapa bagian (chunk) untuk pengolahan paralel
	chunkSize := len(text) / 4
	chunks := make([]string, 4)

	for i := 0; i < 4; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == 3 {
			end = len(text)
		}
		chunks[i] = text[start:end]
	}

	var wg sync.WaitGroup
	ch := make(chan map[rune]int)

	for _, chunk := range chunks {
		wg.Add(1)
		go countLetters(chunk, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := make(map[rune]int)
	for counts := range ch {
		for char, count := range counts {
			result[char] += count
		}
	}

	// Mencetak frekuensi huruf
	for char, count := range result {
		fmt.Printf("%c: %d\n", char, count)
	}
}
