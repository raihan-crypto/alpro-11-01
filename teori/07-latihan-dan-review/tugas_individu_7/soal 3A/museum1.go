package main

import "fmt"

func main() {
	var pengunjung int
	var totalPengunjung, hari, kenaikan int
	prev := -1

	for {
		if _, err := fmt.Scan(&pengunjung); err != nil {
			break
		}

		// Masukan berakhir jika data pengunjung tidak valid
		if pengunjung < 0 || pengunjung > 200 {
			break
		}

		// Hitung kenaikan dibanding hari sebelumnya
		if prev != -1 && pengunjung > prev {
			kenaikan++
		}
		prev = pengunjung

		totalPengunjung += pengunjung
		hari++
	}

	if hari == 0 {
		fmt.Println("0 0")
	} else {
		rataRata := float64(totalPengunjung) / float64(hari)
		fmt.Printf("%d %.2f\n", kenaikan, rataRata)
	}
}

