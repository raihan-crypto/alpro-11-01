package main

import "fmt"

// konversi mengonversi suhu Celsius ke Réaumur, Fahrenheit, dan Kelvin
func konversi(c float64) (float64, float64, float64) {
	r := 0.8 * c
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
