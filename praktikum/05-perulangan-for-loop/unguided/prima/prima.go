package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlahFaktor := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			if jumlahFaktor > 0 {
				fmt.Print(" ")
			}
			fmt.Print(i)
			jumlahFaktor++
		}
	}
	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
