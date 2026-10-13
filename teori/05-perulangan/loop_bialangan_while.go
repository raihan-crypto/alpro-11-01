package main

import "fmt"

func main() {
	var i, n int
	fmt.Scan(&n)
	i = 0
	for i != n {
		i = i + 1
		fmt.Print(i)
	}
}

