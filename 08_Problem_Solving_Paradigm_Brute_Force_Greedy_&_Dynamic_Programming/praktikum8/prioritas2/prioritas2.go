package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	if n <= 1 {
		return 0
	}

	// Membuat slice untuk menyimpan biaya minimum untuk mencapai setiap batu
	cost := make([]int, n)

	// Inisialisasi biaya untuk batu pertama
	cost[0] = 0

	// Inisialisasi biaya untuk batu kedua
	cost[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	// Menghitung biaya minimum untuk batu-batu selanjutnya
	for i := 2; i < n; i++ {
		cost[i] = int(math.Min(
			float64(cost[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))),
			float64(cost[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2])))),
		))
	}

	// Biaya minimum untuk mencapai batu terakhir adalah hasil akhir
	return cost[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
