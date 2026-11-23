package main

import "fmt"

func cekSama(hasil *bool, a, b int) {
	*hasil = a == b
	fmt.Print(*hasil)
}

func main() {
	var a, b int
	var hasil bool
	a = 11
	b = 11
	cekSama(&hasil, a, b)
	fmt.Println(hasil)
}