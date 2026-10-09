package main

import "fmt"

func main() {
	var usia int
	var gaji int

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)
	fmt.Print("Masukkan gaji tahunan (juta): ")
	fmt.Scan(&gaji)

	switch {
	case usia < 18:
		fmt.Println("Masih sekolah")
	case usia >= 18 && usia <= 25 && gaji >= 50:
		fmt.Println("Muda sukses")
	case usia >= 18 && usia <= 25 && gaji < 50:
		fmt.Println("Masih belajar hidup")
	case usia >= 26 && usia <= 40 && gaji >= 100:
		fmt.Println("Pekerja mapan")
	case usia >= 26 && usia <= 40 && gaji < 100:
		fmt.Println("Perlu perbaikan karier")
	case usia > 40 && gaji >= 150:
		fmt.Println("Profesional berpengalaman")
	case usia > 40 && gaji < 150:
		fmt.Println("Perlu evaluasi finansial")
	}
}
