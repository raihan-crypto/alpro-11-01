package main

import "fmt"

const NMAX = 10000

type Arr = [NMAX]int

// procedure bacaData(in/out T : Arr, in n : integer)
// IS : n terdefinisi, T belum berisi data
// FS : T[0..n-1] berisi n nomor sapi hasil input
func bacaData(T *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

// procedure geserKanan(in/out T : Arr, in n : integer)
// IS : T[0..n-1] berisi susunan sapi saat ini
// FS : setiap sapi bergeser satu kandang ke kanan dan
// sapi di kandang terakhir pindah ke kandang pertama
func geserKanan(T *Arr, n int) {
	temp := T[n-1]
	for i := n - 1; i > 0; i-- {
		T[i] = T[i-1]
	}
	T[0] = temp
}

// procedure cetak(in T : Arr, in n : integer)
// IS : T[0..n-1] terdefinisi
// FS : isi T tercetak dalam satu baris dipisahkan spasi
func cetak(T Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var n, k int
	var T Arr

	fmt.Scan(&n, &k)
	bacaData(&T, n)

	// Nilai K dapat mencapai 10^9. Pergeseran sebanyak N kali akan mengembalikan
	// posisi sapi ke susunan awal. Oleh karena itu, cukup dilakukan pergeseran
	// sebanyak (K % N) kali agar program berjalan efisien (O(N) bukannya O(K*N)).
	rotasi := k % n
	for i := 0; i < rotasi; i++ {
		geserKanan(&T, n)
	}

	cetak(T, n)
}
