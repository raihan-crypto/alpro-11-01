package main

import "fmt"

func main() {
	var berat, tinggiCm float64
	fmt.Scan(&berat, &tinggiCm)

	tinggiM := tinggiCm / 100.0
	bmi := berat / (tinggiM * tinggiM)

	var kategori string
	if bmi < 18.5 {
		kategori = "Kurus"
	} else if bmi < 25.0 {
		kategori = "Normal"
	} else if bmi < 30.0 {
		kategori = "Gemuk"
	} else {
		kategori = "Obesitas"
	}

	fmt.Printf("%.2f\n", bmi)
	fmt.Println(kategori)
}

