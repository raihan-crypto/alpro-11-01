package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan dua bilangan (a b): ")
	fmt.Scan(&a, &b)

	fmt.Printf("a > b  : %v\n", a > b)
	fmt.Printf("a == b : %v\n", a == b)
	fmt.Printf("a < b  : %v\n", a < b)
}
