package main

import "fmt"

func main() {
	gold := 0
	platinum := 0
	silver := 0
	totalPoin := 0

	// Berakhir apabila total poin keseluruhan pengguna telah mencapai 500
	for totalPoin < 500 {
		var poin int
		fmt.Scan(&poin)

		// Validasi: masukan harus >= 50
		for poin < 50 {
			fmt.Scan(&poin)
		}

		totalPoin += poin

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

