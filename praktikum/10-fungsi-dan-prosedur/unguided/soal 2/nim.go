package main

import "fmt"

const NMAX = 2002

type tabNIM = [NMAX]int

// procedure cariNIM(in T : tabNIM, in n, x : integer, out idx : integer)
// IS : T[0..n-1] dan x terdefinisi
// FS : idx berisi indeks NIM x pada T, atau -1 jika tidak ditemukan
func cariNIM(T tabNIM, n, x int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i] == x {
			*idx = i
			return
		}
	}
}

// procedure ubahNIM(in/out T : tabNIM, in idx, baru : integer)
// IS : T terdefinisi dan idx adalah indeks yang valid
// FS : T[idx] berisi NIM baru
func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}

// procedure cetak(in T : tabNIM, in n : integer)
// IS : T[0..n-1] terdefinisi
// FS : setiap NIM pada T tercetak satu per baris
func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var n int
	var T tabNIM
	var cari, baru, idx int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	fmt.Scan(&cari)
	cariNIM(T, n, cari, &idx)

	if idx != -1 {
		fmt.Scan(&baru)
		ubahNIM(&T, idx, baru)
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", cari, idx+1)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}
