package main

import "fmt"

// tarifParkir menghitung total tarif parkir berdasarkan lama parkir dalam jam.
// Dua jam pertama: Rp5.000. Setiap jam berikutnya: +Rp3.000/jam.
func tarifParkir(jam int) int {
	if jam <= 2 {
		return 5000
	}
	return 5000 + (jam-2)*3000
}

func main() {
	var jam int
	fmt.Scan(&jam)
	fmt.Println(tarifParkir(jam))
}
