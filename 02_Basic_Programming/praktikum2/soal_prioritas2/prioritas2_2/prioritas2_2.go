package main

import "fmt"

func main() {
    var angka int
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&angka)
	
	fmt.Printf("Faktor dari %d adalah: ", angka)
    for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Print(i, " ")
        }
    }
}