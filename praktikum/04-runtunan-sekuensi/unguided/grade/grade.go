package main

import "fmt"

func main() {
	var nama1, nama2 string
	var nilai int

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&nama1, &nama2)
	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scan(&nilai)

	var grade string

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80:
		grade = "B"
	case nilai >= 70:
		grade = "C"
	case nilai >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Println(nama1, nama2, "mendapatkan nilai", grade)
}
