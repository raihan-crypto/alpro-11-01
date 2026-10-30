package main

import "fmt"

func main() {
	var kelas string
	var usia int
	var hari string

	fmt.Scan(&kelas, &usia, &hari)

	var hargaDasar float64
	switch kelas {
	case "Ekonomi":
		hargaDasar = 100000
	case "Bisnis":
		hargaDasar = 200000
	case "Eksekutif":
		hargaDasar = 350000
	default:
		fmt.Println("Kelas tidak valid")
		return
	}

	harga := hargaDasar
	if hari == "Sabtu" || hari == "Minggu" {
		harga += harga * 0.20
	}

	if usia < 3 {
		harga = 0
	} else if usia >= 3 && usia <= 12 {
		harga -= harga * 0.50
	} else if usia >= 60 {
		harga -= harga * 0.30
	}

	fmt.Println(int(harga))
}

