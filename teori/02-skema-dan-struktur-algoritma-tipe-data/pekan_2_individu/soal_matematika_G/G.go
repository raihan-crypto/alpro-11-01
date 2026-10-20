package main

import "fmt"

func main() {
	var p, l int
	var luas, keliling int

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&l)

	luas = p * l
	keliling = 2 * (p + l)

	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)
}
