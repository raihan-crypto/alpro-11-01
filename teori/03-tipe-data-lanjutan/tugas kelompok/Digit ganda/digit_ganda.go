package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	digitPertama := n / 10   // digit puluhan, contoh: 15 -> 1
	digitKedua := n % 10     // digit satuan,  contoh: 15 -> 5

	// Susun menjadi 4 digit: d1 d1 d2 d2
	// contoh: 1 1 5 5 -> 1*1000 + 1*100 + 5*10 + 5
	hasil := digitPertama*1000 + digitPertama*100 + digitKedua*10 + digitKedua

	fmt.Println(hasil)
}
