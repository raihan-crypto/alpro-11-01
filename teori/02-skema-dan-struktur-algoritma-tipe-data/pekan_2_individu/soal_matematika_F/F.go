package main

import "fmt"

func main() {
	var x float64
	var hasil float64

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	hasil = (x*x + 2*x + 1) / (x - 3)

	fmt.Println(hasil)
}
