package main

import "fmt"

// faktorial menghitung n!
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// kombinasi menghitung C(n, r) = n! / (r! * (n-r)!)
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fx := faktorial(x)
	fy := faktorial(y)

	var c int
	if x >= y {
		c = kombinasi(x, y)
	} else {
		c = kombinasi(y, x)
	}

	fmt.Println(fx, fy, c)
}
