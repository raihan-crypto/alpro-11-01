package main

import "fmt"

func main() {
	total := 0

	for hari := 1; hari <= 5; hari++ {
		var pengunjung int
		for {
			fmt.Printf("Hari %d : ", hari)
			fmt.Scan(&pengunjung)

			if pengunjung >= 0 && pengunjung <= 200 {
				total += pengunjung
				break
			}
		}
	}

	fmt.Printf("Jumlah pengunjung : %d orang\n", total)
}

