package main

import "fmt"

func main() {
	var opsi int

	fmt.Println("============== Menu ==============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opsi)

	if opsi == 1 {
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
	} else if opsi == 0 {
		fmt.Println("Keluar dari program.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
