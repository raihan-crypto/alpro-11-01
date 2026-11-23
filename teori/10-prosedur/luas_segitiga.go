package main

import "fmt"

func hitungLuas(alas, tinggi int, luas *float64) {
	*luas = float64(alas) * float64(tinggi) / 2.0
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a, t int
		var luas float64

		fmt.Scan(&a, &t)
		hitungLuas(a, t, &luas)
		fmt.Println(luas)
	}
}
