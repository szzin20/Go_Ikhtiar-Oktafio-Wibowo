package main

import "fmt"

func main() {
    var nilai int
    fmt.Print("Masukkan nilai: ")
    fmt.Scan(&nilai)

    var grade string

    if nilai >= 80 && nilai <= 100 {
        grade = "A"
    } else if nilai >= 65 && nilai <= 79 {
        grade = "B"
    } else if nilai >= 50 && nilai <= 64 {
        grade = "C"
    } else if nilai >= 35 && nilai <= 49 {
        grade = "D"
    } else if nilai >= 0 && nilai <= 34 {
        grade = "E"
    } else {
        grade = "Nilai Invalid"
    }

    fmt.Println("Grade : ", grade)
}