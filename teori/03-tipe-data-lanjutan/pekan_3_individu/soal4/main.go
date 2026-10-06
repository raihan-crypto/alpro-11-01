package main

import "fmt"

func main() {
	var x, low, high int
	fmt.Print("Masukkan tiga bilangan (x low high): ")
	fmt.Scan(&x, &low, &high)

	fmt.Printf("Dalam rentang [%d, %d] : %v\n", low, high, x >= low && x <= high)
}
