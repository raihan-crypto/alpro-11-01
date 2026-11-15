package main

import "fmt"

// denom menghitung jumlah lembar pecahan Rp10.000, Rp2.000, dan Rp1.000
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
