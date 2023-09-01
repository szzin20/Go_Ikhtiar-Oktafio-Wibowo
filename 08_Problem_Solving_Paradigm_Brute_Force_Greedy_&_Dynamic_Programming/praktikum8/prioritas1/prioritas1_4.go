package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	var x, y, z int

	for x = 1; x <= 10000; x++ {
		for y = 1; y <= 10000; y++ {
			z = a - x - y
			if z > 0 && x*y*z == b && x*x+y*y+z*z == c {
				break
			}
		}
		if y <= 10000 {
			break
		}
	}

	if x <= 10000 && y <= 10000 {
		fmt.Printf("%d %d %d\n", x, y, z)
	} else {
		fmt.Println("No solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // No solution
	SimpleEquations(6, 6, 14) // 1 2 3
}