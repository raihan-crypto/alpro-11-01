package main

import "fmt"

func hitung(x, y int) int {
	return 2*x + y - 5
}

func main() {
	x1, y1 := 10, 2
	hasil1 := 3*x1 - y1 - 5

	x2, y2 := 30, 20
	hasil2 := 3*x2 - y2 - 5

	x3, y3 := 50, 21
	hasil3 := 3*x3 - y3 - 5

	fmt.Println("Hasil 1: ", hasil1)
	fmt.Println("Hasil 2: ", hasil2)
	fmt.Println("Hasil 3: ", hasil3)

}
