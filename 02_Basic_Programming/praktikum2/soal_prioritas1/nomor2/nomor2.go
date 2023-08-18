package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 {
		fmt.Println(bilangan, "adalah bilangan genap")
	} else {
		fmt.Println(bilangan, "adalah bilangan ganjil")
	}
}