package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	limit := n
	if m < limit {
		limit = m
	}

	fpb := 1
	for i := limit; i >= 1; i-- {
		if n%i == 0 && m%i == 0 {
			fpb = i
			break
		}
	}

	fmt.Println(fpb)
}

