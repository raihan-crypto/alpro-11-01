package main

import "fmt"

func main() {
	var belanja, potongan float64
	var asesmen bool

	fmt.Scan(&belanja, &asesmen)
	if asesmen {
		potongan = belanja * 0.35
		belanja = belanja - potongan
	}
	fmt.Println(belanja)
}