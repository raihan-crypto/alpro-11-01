package main

import "fmt"

const NMAX = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa = [NMAX]gempa

// function kategori(m : real) -> string
func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m < 6.0 {
		return "Sedang"
	} else if m < 7.0 {
		return "Kuat"
	} else {
		return "Besar"
	}
}

// function rataMagnitudo(T : tabGempa, n : integer) -> real
func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}
	return total / float64(n)
}

// function idxTerkuat(T : tabGempa, n : integer) -> integer
// jika ada lebih dari satu gempa dengan magnitudo terbesar,
// gempa terkuat adalah yang pertama muncul
func idxTerkuat(T tabGempa, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[idxMax].magnitudo {
			idxMax = i
		}
	}
	return idxMax
}

// function hitungKategori(T : tabGempa, n : integer, k : string) -> integer
// fungsi hitungKategori wajib memanggil fungsi kategori di dalamnya
func hitungKategori(T tabGempa, n int, k string) int {
	count := 0
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var T tabGempa

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n", T[i].nama, T[i].tahun, T[i].magnitudo, kategori(T[i].magnitudo))
	}

	terkuat := idxTerkuat(T, n)
	fmt.Printf("Rata-rata magnitudo: %.2f\n", rataMagnitudo(T, n))
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", T[terkuat].nama, T[terkuat].tahun, T[terkuat].magnitudo)
	fmt.Printf("Jumlah gempa kategori Besar: %d\n", hitungKategori(T, n, "Besar"))
}
