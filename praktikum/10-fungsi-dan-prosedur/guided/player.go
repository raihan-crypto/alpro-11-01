package main

import "fmt"

const NMAX = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer = [NMAX]player

// function rataRata(T : tabPlayer, n : integer) -> real
// mengembalikan rata-rata skor dari n player
func rataRata(T tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].score
	}
	return float64(total) / float64(n)
}

// function idxTertinggi(T : tabPlayer, n : integer) -> integer
// mengembalikan indeks player dengan skor tertinggi;
// jika ada skor sama, ambil yang pertama muncul
func idxTertinggi(T tabPlayer, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].score > T[idxMax].score {
			idxMax = i
		}
	}
	return idxMax
}

// function hitungDiAtas(T : tabPlayer, n : integer, batas : real) -> integer
// mengembalikan banyaknya player dengan skor lebih dari batas
func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var T tabPlayer

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	rata := rataRata(T, n)
	terbaik := idxTertinggi(T, n)
	diAtas := hitungDiAtas(T, n, rata)

	fmt.Printf("Rata-rata skor: %.2f\n", rata)
	fmt.Printf("Player terbaik: %s (%d)\n", T[terbaik].nickname, T[terbaik].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", diAtas)
}
