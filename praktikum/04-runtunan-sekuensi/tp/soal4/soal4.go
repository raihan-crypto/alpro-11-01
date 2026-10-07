package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai ujian (0-100): ")
	fmt.Scan(&nilai)

	var grade string

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80:
		grade = "B"
	case nilai >= 70:
		grade = "C"
	case nilai >= 60:
		grade = "D"
	case nilai >= 0:
		grade = "E"
	default:
		fmt.Println("Nilai tidak valid!")
		return
	}

	fmt.Printf("Nilai %d mendapat grade: %s\n", nilai, grade)
}
