package main

import "fmt"

func main() {
	var nama1 string
	var nilai int

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&nama1)
	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scan(&nilai)

	var grade string

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 {
		grade = "B"
	} else if nilai >= 70 {
		grade = "C"
	} else if nilai >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Println(nama1, "mendapatkan nilai", grade)
}
