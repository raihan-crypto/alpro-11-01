package main

import "fmt"

func main() {
	var harga, persen float64
	var potongan, hargaAkhir float64

	fmt.Print("Masukkan harga: ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan persenan: ")
	fmt.Scan(&persen)

	potongan = harga * persen / 100
	hargaAkhir = (harga) - potongan

	fmt.Println("Harga Akhir:", hargaAkhir)
}