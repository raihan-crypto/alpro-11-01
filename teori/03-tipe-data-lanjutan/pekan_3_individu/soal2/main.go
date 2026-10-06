package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan (n): ")
	fmt.Scan(&n)

	fmt.Printf("Genap (n %% 2 == 0) : %v\n", n%2 == 0)
}
