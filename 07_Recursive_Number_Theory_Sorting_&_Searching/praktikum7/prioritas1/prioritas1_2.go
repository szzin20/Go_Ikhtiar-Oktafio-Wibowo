package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	countMap := make(map[string]int)

	// Menghitung kemunculan setiap barang
	for _, item := range items {
		countMap[item]++
	}

	// Mengonversi map ke dalam slice pair
	var pairs []pair
	for name, count := range countMap {
		pairs = append(pairs, pair{name, count})
	}

	// Fungsi untuk mengurutkan slice pairs berdasarkan count
	sortPairsByCount(pairs)

	return pairs
}

func sortPairsByCount(pairs []pair) {
	for i := 0; i < len(pairs)-1; i++ {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[j].count > pairs[i].count {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})
	for _, list := range pairs {
		fmt.Println(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, list := range pairs {
		fmt.Println(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"})
	for _, list := range pairs {
		fmt.Println(list.name, " -> ", list.count, " ")
	}
}