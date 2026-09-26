package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("=== Program Kalkulator Sederhana ===")
	fmt.Print("Masukkan bilangan bulat pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan bulat kedua   : ")
	fmt.Scan(&b)

	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	mod := a % b

	fmt.Println("\n=== Hasil Perhitungan ===")
	fmt.Printf("Penjumlahan (%d + %d) = %d\n", a, b, tambah)
	fmt.Printf("Pengurangan (%d - %d) = %d\n", a, b, kurang)
	fmt.Printf("Perkalian   (%d * %d) = %d\n", a, b, kali)
	fmt.Printf("Pembagian   (%d / %d) = %d\n", a, b, bagi)
	fmt.Printf("Sisa Bagi   (%d %% %d) = %d\n", a, b, mod)
}