package main

import "fmt"

func denom(n int) (int, int, int) {
	lembar10000 := n / 10000
	sisa := n % 10000
	lembar2000 := sisa / 2000
	sisa = sisa % 2000
	lembar1000 := sisa / 1000
	return lembar10000, lembar2000, lembar1000
}

func main() {
	var n int
	fmt.Scan(&n)
	a, b, c := denom(n)
	fmt.Println(a, b, c)
}
