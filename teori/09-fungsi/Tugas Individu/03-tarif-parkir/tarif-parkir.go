package main

import "fmt"

func tarifParkir(jam int) int {
	if jam <= 2 {
		return 5000
	}
	return 5000 + (jam-2)*3000
}

func main() {
	var jam int
	fmt.Scan(&jam)
	fmt.Println(tarifParkir(jam))
}
