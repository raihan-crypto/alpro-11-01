package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	max := n
	if m > max {
		max = m
	}

	kpk := max
	for kpk%n != 0 || kpk%m != 0 {
		kpk += max
	}

	fmt.Println(kpk)
}

