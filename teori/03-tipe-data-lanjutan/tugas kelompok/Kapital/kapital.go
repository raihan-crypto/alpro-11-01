package main

import "fmt"

func main() {
	var c string

	fmt.Scan(&c)

	// Ambil karakter pertama dari input
	karakter := rune(c[0])

	// Huruf kapital berada di rentang 'A' sampai 'Z'
	isKapital := karakter >= 'A' && karakter <= 'Z'

	fmt.Println(isKapital)
}
