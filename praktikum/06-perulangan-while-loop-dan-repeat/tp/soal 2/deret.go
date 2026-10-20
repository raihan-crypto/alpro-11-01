package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := 0
	i := 1
	for i <= n {
		s += i
		i++
	}

	fmt.Println(s)
}
