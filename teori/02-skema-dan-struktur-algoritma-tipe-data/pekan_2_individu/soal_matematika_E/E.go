package main

import "fmt"

func main() {
	var x, y int
	var hasil float64

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	hasil = 5*float64(x*x) - 2*float64(x*y) + float64(y*y*y)/float64(x+1)

	fmt.Println("Hasil: ", hasil)
}
