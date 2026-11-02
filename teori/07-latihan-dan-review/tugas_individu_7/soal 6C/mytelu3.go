package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	gold := 0
	platinum := 0
	silver := 0

	for i := 0; i < n; i++ {
		var poin int
		fmt.Scan(&poin)

		// Validasi: masukan harus >= 50
		for poin < 50 {
			fmt.Scan(&poin)
		}

		if poin > 200 {
			gold++
		} else if poin >= 100 {
			platinum++
		} else {
			silver++
		}
	}

	fmt.Println(gold, platinum, silver)
}

