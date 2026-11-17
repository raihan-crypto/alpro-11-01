package main

import "fmt"

func tambah(x, y int) int {
	x = x + y
	return x
}

func main() {
	a := 4
	b := tambah(a, 2)
	fmt.Println(a, b)

	a = tambah(b, 3)
	fmt.Println(a, b)
}