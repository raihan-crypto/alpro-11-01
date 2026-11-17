package main

import "fmt"

func deretKuadrat(n int) int {
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return total
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(deretKuadrat(n))
}
