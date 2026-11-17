package main

import "fmt"

func kuadran(x, y int) string {
		return "Titik Asal"
	} else if x == 0 {
		return "Sumbu Y"
	} else if y == 0 {
		return "Sumbu X"
	} else if x > 0 && y > 0 {
		return "Kuadran I"
	} else if x < 0 && y > 0 {
		return "Kuadran II"
	} else if x < 0 && y < 0 {
		return "Kuadran III"
	} else {
		return "Kuadran IV"
	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(kuadran(x, y))
}
