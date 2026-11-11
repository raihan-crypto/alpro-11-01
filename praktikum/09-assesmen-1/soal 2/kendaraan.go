package main

import "fmt"

func main() {
	var kendaraan string
	var motorCount int
	var mobilCount int
	var trukCount int

	for {
		fmt.Scan(&kendaraan)

		if kendaraan != "M" && kendaraan != "K" && kendaraan != "T" {
			break
		}

		if kendaraan == "M" {
			motorCount = motorCount + 1
		} else if kendaraan == "K" {
			mobilCount = mobilCount + 1
		} else if kendaraan == "T" {
			trukCount = trukCount + 1
		}
	}

	fmt.Printf("Tipe M = %d\n", motorCount)
	fmt.Printf("Tipe K = %d\n", mobilCount)
	fmt.Printf("Tipe T = %d\n", trukCount)
}
