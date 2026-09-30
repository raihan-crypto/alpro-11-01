package main

import "fmt"

func main() {
	var uang int

	fmt.Println("=== Program Pemecah Pecahan Uang ===")
	fmt.Print("Masukkan nominal uang: ")
	fmt.Scan(&uang)

	totalAwal := uang

	sepuluhRibu := uang / 10000
	uang = uang % 10000

	limaRibu := uang / 5000
	uang = uang % 5000

	seribu := uang / 1000

	fmt.Printf("\nHasil Pemecahan Uang Rp %d:\n", totalAwal)
	fmt.Printf("- Pecahan 10.000 : %d lembar\n", sepuluhRibu)
	fmt.Printf("- Pecahan 5.000  : %d lembar\n", limaRibu)
	fmt.Printf("- Pecahan 1.000  : %d lembar\n", seribu)
}