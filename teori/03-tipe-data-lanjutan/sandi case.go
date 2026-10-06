package main

import "fmt"

func main() {
	var jumlah, n, depan, belakang, total int
	fmt.Scan(&jumlah)

for i := 0; i < jumlah; i++ {
	fmt.Scan(&n)
	depan = n / 1000
	belakang = n % 10
	total += depan + belakang
}
fmt.Println(total)

}