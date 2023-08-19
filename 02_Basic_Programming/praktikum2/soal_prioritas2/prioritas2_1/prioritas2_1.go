package main

import "fmt"

func main() {
    rows := 5 
    for i := 1; i <= rows; i++ {
      
        for j := 1; j <= rows-i; j++ {
            fmt.Print(" ")
        }

        for k := 1; k <=1*i; k++ {
            fmt.Print("* ")
        }

        fmt.Println() 
    }
}