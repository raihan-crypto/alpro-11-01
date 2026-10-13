package main 

import "fmt"

func main() {
	var n int
	jumlah := 0

	fmt.Scan(&n)
	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ") 
		jumlah += i 
	}
	fmt.Println()
	fmt.Println(jumlah)
}