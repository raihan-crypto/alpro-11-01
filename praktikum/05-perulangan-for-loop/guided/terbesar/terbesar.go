package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	total := 0
	tabungan := a

	for i := 1; i <= n; i++ {
		total = total + tabungan
		tabungan = tabungan + b
	}

	fmt.Println(total)
}
