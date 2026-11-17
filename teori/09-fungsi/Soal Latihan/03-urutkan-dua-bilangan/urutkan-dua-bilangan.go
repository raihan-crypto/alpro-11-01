package main

import "fmt"

func urutDua(a, b int) (int, int) {
		return a, b
	}
	return b, a
}

func main() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil || a == b {
			break
		}
		kecil, besar := urutDua(a, b)
		fmt.Println(kecil, besar)
	}
}
