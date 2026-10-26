package main

import "fmt"

func main() {
	var t1, t2, t3, t4, t5 float64
	fmt.Scan(&t1, &t2, &t3, &t4, &t5)

	if t1 < t2 && t2 < t3 && t3 < t4 && t4 < t5 {
		fmt.Println("Stabil naik")
	} else if t1 > t2 && t2 > t3 && t3 > t4 && t4 > t5 {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}
