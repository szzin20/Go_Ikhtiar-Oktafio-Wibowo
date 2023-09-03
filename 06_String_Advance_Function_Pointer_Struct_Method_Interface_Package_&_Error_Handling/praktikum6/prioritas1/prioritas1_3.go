package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var result string

	// Loop through the shorter string
	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			substring := a[i:j]
			if len(substring) > len(result) && isSubstring(substring, b) {
				result = substring
			}
		}
	}

	return result
}

func isSubstring(substring, str string) bool {
	return len(substring) <= len(str) && str[:len(substring)] == substring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
