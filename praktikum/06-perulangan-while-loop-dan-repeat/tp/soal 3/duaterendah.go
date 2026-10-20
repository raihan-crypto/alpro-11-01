package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var namaTerendah, namaKedua string
	var nilaiTerendah, nilaiKedua int

	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if i == 0 {
			namaTerendah = nama
			nilaiTerendah = nilai
		} else if i == 1 {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah
				namaTerendah = nama
				nilaiTerendah = nilai
			} else {
				namaKedua = nama
				nilaiKedua = nilai
			}
		} else {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah
				namaTerendah = nama
				nilaiTerendah = nilai
			} else if nilai < nilaiKedua {
				namaKedua = nama
				nilaiKedua = nilai
			}
		}
		i++
	}

	fmt.Println(namaTerendah)
	fmt.Println(namaKedua)
}
