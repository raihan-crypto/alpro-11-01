package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scan(&celsius)

	reamur := (4.0 / 5.0) * celsius

	fmt.Printf("Suhu dalam derajat Reamur: %.2f\n", reamur)
}
