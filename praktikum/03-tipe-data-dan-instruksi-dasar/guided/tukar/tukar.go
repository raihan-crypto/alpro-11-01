package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scanln(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}