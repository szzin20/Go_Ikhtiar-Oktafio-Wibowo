//luas trapesium

package main

import "fmt"

func main() {
	var sisiA int 
	fmt.Print("Masukkan sisi A : ")
	fmt.Scan(&sisiA)
	
	var sisiB int 
	fmt.Print("Masukkan sisi B : ")
	fmt.Scan(&sisiB)
	
	var tinggi int 
	fmt.Print("Masukkan tinggi trapesium : ")
	fmt.Scan(&tinggi)

	//sisiB := 6
	//tinggi := 4
	luas := 0.5 * (float64(sisiA)+ float64(sisiB)) * float64(tinggi)
	fmt.Println("Luas trapesium:", luas)
}