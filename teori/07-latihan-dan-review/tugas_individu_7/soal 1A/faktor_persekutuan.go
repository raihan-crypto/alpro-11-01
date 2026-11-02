package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	limit := n
	if m < limit {
		limit = m
	}

	first := true
	for i := 1; i <= limit; i++ {
		if n%i == 0 && m%i == 0 {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(i)
			first = false
		}
	}
	fmt.Println()
}

