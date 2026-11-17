package main

import "fmt"

func konversi(c float64) (float64, float64, float64) {
	f := 1.8*c + 32
	k := c + 273.15
	return r, f, k
}

func main() {
	var c float64
	fmt.Scan(&c)
	r, f, k := konversi(c)
	fmt.Printf("%.2f %.2f %.2f\n", r, f, k)
}
