package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	var count int
	count = 0

	var i int
	for i = l; i <= r; i++ {
		var angka int
		angka = i

		var digitSebelumnya int
		digitSebelumnya = angka % 10
		angka = angka / 10

		var nonDecreasing bool
		nonDecreasing = true

		for angka > 0 {
			var digitSaatIni int
			digitSaatIni = angka % 10

			if digitSaatIni > digitSebelumnya {
				nonDecreasing = false
				break
			}

			digitSebelumnya = digitSaatIni
			angka = angka / 10
		}

		if !nonDecreasing {
			continue
		}

		angka = i
		var jumlahDigit int
		jumlahDigit = 0

		for angka > 0 {
			jumlahDigit = jumlahDigit + (angka % 10)
			angka = angka / 10
		}

		if i == 0 {
			jumlahDigit = 0
		}

		var prima bool
		prima = true

		if jumlahDigit < 2 {
			prima = false
		} else {
			var j int
			for j = 2; j*j <= jumlahDigit; j++ {
				if jumlahDigit%j == 0 {
					prima = false
					break
				}
			}
		}

		if prima {
			count = count + 1
		}
	}

	fmt.Println(count)
}
