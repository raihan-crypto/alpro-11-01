package main

import "fmt"

func main() {
	var hari int

	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&hari)

	tahun := hari / 360
	hari = hari % 360

	bulan := hari / 30
	hari = hari % 30

	minggu := hari / 7
	sisaHari := hari % 7

	fmt.Printf("Tahun  : %d\n", tahun)
	fmt.Printf("Bulan  : %d\n", bulan)
	fmt.Printf("Minggu : %d\n", minggu)
	fmt.Printf("Hari   : %d\n", sisaHari)
}
