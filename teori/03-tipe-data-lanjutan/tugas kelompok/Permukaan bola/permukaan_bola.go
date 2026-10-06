package main

import "fmt"

func main() {
	var r float64

	fmt.Scan(&r)

	pi := 22.0 / 7.0
	luas := 4 * pi * r * r

	fmt.Println(luas)
}
