package main

import "fmt"

func main() {
	var poin int
	fmt.Scan(&poin)

	// Validasi: masukan harus >= 50
	for poin < 50 {
		fmt.Scan(&poin)
	}

	if poin > 200 {
		fmt.Println("Gold user")
	} else if poin >= 100 {
		fmt.Println("Platinum user")
	} else {
		fmt.Println("Silver user")
	}
}

