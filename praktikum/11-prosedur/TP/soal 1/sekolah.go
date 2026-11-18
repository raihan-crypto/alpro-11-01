package main

import "fmt"

const NMAX = 1001

type tabMurid = [NMAX]string

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false
	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1
			return
		}
	}
}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			*jumlah++
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabMurid
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	var nama string
	fmt.Scan(&nama)

	var posisi int
	var ketemu bool
	cariMurid(T, n, nama, &posisi, &ketemu)

	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}

	var jumlah int
	hitungHuruf(T, n, nama[0], &jumlah)
	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", nama[0], jumlah)
}
