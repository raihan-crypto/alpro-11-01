package main

import "fmt"

func konversi(km, m, cm int) int {
	return km*100000 + m*100 + cm
}

func main() {
	var km, m, cm int
	fmt.Scan(&km, &m, &cm)
	hasil := konversi(km, m, cm)
	fmt.Printf("Hasil konversi = %d cm\n", hasil)
}
