package main

import "fmt"

func konversi(km, m, cm int) int {
}

func main() {
	var km, m, cm int
	fmt.Scan(&km, &m, &cm)
	hasil := konversi(km, m, cm)
	fmt.Printf("Hasil konversi = %d cm\n", hasil)
}
