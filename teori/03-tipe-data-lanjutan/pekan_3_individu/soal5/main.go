package main

import "fmt"

func main() {
	var p, q int
	fmt.Print("Masukkan dua bilangan (p q): ")
	fmt.Scan(&p, &q)

	fmt.Printf("p genap OR q genap    : %v\n", (p%2 == 0) || (q%2 == 0))
	fmt.Printf("p ganjil AND q ganjil : %v\n", (p%2 != 0) && (q%2 != 0))
	fmt.Printf("NOT(p == q)           : %v\n", !(p == q))
}
