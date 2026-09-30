package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukan suhu celcius: ")
	fmt.Scanln(&celcius)

	fmt.Print("Suhu dalam kelvin: ")
	fmt.Println(celcius + 273)

}