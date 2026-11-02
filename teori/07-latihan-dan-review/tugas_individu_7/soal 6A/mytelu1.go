package main

import "fmt"

func main() {
	var poin int
	fmt.Scan(&poin)

	if poin > 200 {
		fmt.Println("Gold user")
	} else if poin >= 100 {
		fmt.Println("Platinum user")
	} else {
		fmt.Println("Silver user")
	}
}

