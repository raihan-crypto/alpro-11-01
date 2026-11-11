package main

import "fmt"

func main() {
	var angka int
	var sebelumnya int
	var tertinggi int
	var terendah int
	var total int
	var jumlah int

	fmt.Scan(&angka)

	tertinggi = angka
	terendah = angka
	total = angka
	jumlah = 1
	sebelumnya = angka

	for {
		fmt.Scan(&angka)

		if angka == sebelumnya {
			break
		}

		if angka > tertinggi {
			tertinggi = angka
		}

		if angka < terendah {
			terendah = angka
		}

		total = total + angka
		jumlah = jumlah + 1
		sebelumnya = angka
	}

	var rata float64
	rata = float64(total) / float64(jumlah)

	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, rata)
}
