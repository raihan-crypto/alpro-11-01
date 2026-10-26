package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan%3 == 0 {
		fmt.Println("Kelipatan 3")
	}
	if bilangan%5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}
