package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	kabisat := (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)

	var jumlahHari int
	valid := true

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		jumlahHari = 31
	case "Apr", "Jun", "Sep", "Nov":
		jumlahHari = 30
	case "Feb":
		if kabisat {
			jumlahHari = 29
		} else {
			jumlahHari = 28
		}
	default:
		valid = false
	}

	if valid {
		fmt.Println(jumlahHari)
	} else {
		fmt.Println("-")
	}
}
