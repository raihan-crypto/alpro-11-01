package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var lebar int
	fmt.Scan(&lebar)

	var tersempit int
	tersempit = lebar

	var i int
	for i = 1; i < n; i++ {
		fmt.Scan(&lebar)
		if lebar < tersempit {
			tersempit = lebar
		}
	}

	fmt.Println(tersempit)
}
