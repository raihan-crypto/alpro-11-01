package main

import "fmt"

func main() {
	var m, nilai, lulus int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		jumlah := 0
		for j := 0; j < 6; j++ {
			fmt.Scan(&nilai)
			jumlah += nilai
		}

		rata := float64(jumlah) / 6
		if rata >= 3.33 {
			fmt.Println("Ya")
			lulus++
		} else {
			fmt.Println("Tidak")
		}
	}
	fmt.Println(lulus)
}