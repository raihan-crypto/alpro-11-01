package main 

import "fmt"

func hitung(x, y int) int {
	return 2*x + y -5
}

func main() {
	fmt.Println(hitung(10, 2)) // a dan 5 ini adalah argumen
	fmt.Println(hitung(10, 20))
	fmt.Println(hitung(50, 21))


}