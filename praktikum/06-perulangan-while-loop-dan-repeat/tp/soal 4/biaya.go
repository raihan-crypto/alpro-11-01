package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0
	tahun := 1

	for tahun < n {
		biaya = biaya * 1.06
		tahun++
	}

	fmt.Printf("%.0f\n", biaya)
}
