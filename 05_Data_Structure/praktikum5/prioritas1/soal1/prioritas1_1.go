package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make([]string, 0, len(arrayA)+len(arrayB))
	seen := make(map[string]bool)

	for _, name := range arrayA {
		if !seen[name] {
			merged = append(merged, name)
			seen[name] = true
		}
	}

	for _, name := range arrayB {
		if !seen[name] {
			merged = append(merged, name)
			seen[name] = true
		}
	}

	return merged
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}