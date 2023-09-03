package main

import "fmt"

func findMaxMin(numbers []int) (int, int) {
    if len(numbers) == 0 {
        return 0, 0
    }

    max := numbers[0]
    min := numbers[0]

    for _, num := range numbers {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }

    return max, min
}

func main() {
    var inputNumbers [6]int

    fmt.Println("Enter 6 numbers:")
    for i := 0; i < 6; i++ {
        fmt.Scan(&inputNumbers[i])
    }

    max, min := findMaxMin(inputNumbers[:])

    fmt.Printf("%d is the maximum number\n", max)
    fmt.Printf("%d is the minimum number\n", min)
}
