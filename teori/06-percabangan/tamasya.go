package main

import "fmt"

func main(){
	var N, mobil int
	fmt.Scan(&N)

	mobil = N / 7
	if N % 7 != 0 {
	mobil = mobil + 1
}
	fmt.Println(mobil)
}